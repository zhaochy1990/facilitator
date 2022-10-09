package agent_manager

import (
	"errors"
	"fmt"
	"github.com/zhaochy1990/facilitator/internal/model"
	"github.com/zhaochy1990/facilitator/internal/persister"
	"net/http"
	"time"
)

type AgentManager struct {
	persister persister.Persister
}

func NewAgentManager(p persister.Persister) *AgentManager {
	return &AgentManager{persister: p}
}

func (a AgentManager) Create(agent model.Agent) error {
	err := a.ping(agent, 3)
	if err != nil {
		return err
	}

	err = a.persister.SaveAgent(agent)
	if err != nil {
		return err
	}
	return nil
}

func (a AgentManager) ping(agent model.Agent, retry int) error {
	client := http.Client{
		Timeout: 5 * time.Second,
	}

	for retry >= 0 {
		response, err := client.Get(agent.Host)
		if err != nil {
			if retry == 0 {
				return err
			} else {
				retry -= 1
				continue
			}
		}
		if response.StatusCode != http.StatusOK {
			if retry == 0 {
				msg := fmt.Sprintf("can not get status ok (200) from agent %s[%s], ", agent.Name, agent.Host)
				return errors.New(msg)
			} else {
				retry -= 1
				continue
			}
		}
		break
	}
	return nil
}
