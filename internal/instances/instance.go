package instances

import (
	"errors"
	"sync"

	"github.com/rs/zerolog/log"
)

type Instance struct {
	id        string
	Reference string
	runFunc   func(stopChan <-chan struct{})
	stop      chan struct{}
	wg        *sync.WaitGroup
	IsRunning bool
	mu        sync.Mutex
}

func (i *Instance) GetID() string {
	return i.id
}

func (i *Instance) Start() error {
	i.mu.Lock()
	defer i.mu.Unlock()

	if i.IsRunning {
		log.Warn().Msgf("INSTANCE: Instancia %s já está em execução.\n", i.id)
		return errors.New("INSTANCE: Instancia " + i.id + " já está em execução")
	}

	i.stop = make(chan struct{})
	i.wg.Add(1)
	go func() {
		defer i.wg.Done()
		i.runFunc(i.stop)
	}()
	i.IsRunning = true
	log.Info().Msgf("INSTANCE: Instancia %s iniciada.\n", i.id)
	return nil
}

func (i *Instance) Stop() error {
	i.mu.Lock()
	defer i.mu.Unlock()

	if !i.IsRunning {
		log.Warn().Msgf("INSTANCE: Instancia %s não está em execução.\n", i.id)
		return errors.New("INSTANCE: Instancia " + i.id + " não está em execução")
	}

	close(i.stop)
	i.IsRunning = false
	log.Info().Msgf("INSTANCE: Instancia %s parada.\n", i.id)
	return nil
}

type InstanceManager struct {
	instances map[string]*Instance
	mu        sync.Mutex
	wg        sync.WaitGroup
}

func NewManager() *InstanceManager {
	return &InstanceManager{
		instances: make(map[string]*Instance),
	}
}

func (m *InstanceManager) GetInstance(id string) (*Instance, error) {
	m.mu.Lock()
	defer m.mu.Unlock()

	instance, exists := m.instances[id]
	if !exists {
		return nil, errors.New("INSTANCE MANAGER: Instancia com ID " + id + " não encontrada")
	}

	return instance, nil
}

func (m *InstanceManager) GetAllInstances() map[string]*Instance {
	m.mu.Lock()
	defer m.mu.Unlock()

	return m.instances
}

func (m *InstanceManager) AddInstance(id string, reference string, runFunc func(stopChan <-chan struct{})) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	if _, exists := m.instances[id]; exists {
		return errors.New("INSTANCE MANAGER: Instancia com ID " + id + " já existe")
	}

	instance := &Instance{
		id:        id,
		Reference: reference,
		runFunc:   runFunc,
		wg:        &m.wg,
	}
	m.instances[id] = instance
	log.Info().Msgf("INSTANCE MANAGER: Instancia %s adicionada.\n", id)
	return nil
}

func (m *InstanceManager) StartInstance(id string) error {
	i, err := m.GetInstance(id)

	if err != nil {
		return err
	}

	return i.Start()
}

func (m *InstanceManager) StopInstance(id string) error {
	m.mu.Lock()
	instance, exists := m.instances[id]
	m.mu.Unlock()

	if !exists {
		return errors.New("INSTANCE MANAGER: Instancia com ID " + id + " não encontrada")
	}

	return instance.Stop()
}

func (m *InstanceManager) DeleteInstance(id string) error {
	m.mu.Lock()
	instance, exists := m.instances[id]
	m.mu.Unlock()

	if !exists {
		return errors.New("INSTANCE MANAGER: Instancia com ID " + id + " não encontrada para deletar")
	}

	instance.mu.Lock()
	IsRunning := instance.IsRunning
	instance.mu.Unlock()

	if IsRunning {
		return errors.New("INSTANCE MANAGER: Instancia " + id + " está em execução. Pare-a antes de deletar.")
	}

	m.mu.Lock()
	delete(m.instances, id)
	m.mu.Unlock()
	log.Info().Msgf("INSTANCE MANAGER: Instancia %s deletada.\n", id)
	return nil
}

func (m *InstanceManager) DeleteAllInstances() error {
	m.mu.Lock()
	defer m.mu.Unlock()

	for id, instance := range m.instances {
		err := instance.Stop()
		if err != nil {
			return err
		}
		delete(m.instances, id)
		log.Info().Msgf("INSTANCE MANAGER: Instancia %s deletada.\n", id)
	}
	return nil
}

func (m *InstanceManager) StopAllInstances() error {
	log.Info().Msgf("INSTANCE MANAGER: Parando todas as instâncias.")
	m.mu.Lock()
	defer m.mu.Unlock()

	var firstErr error
	for id, instance := range m.instances {
		err := instance.Stop()
		if err != nil {
			if firstErr == nil {
				firstErr = err
			}
		} else {
			log.Info().Msgf("INSTANCE MANAGER: Instancia %s parada.\n", id)
		}
	}
	return firstErr
}

func (m *InstanceManager) StartAllInstances() error {
	m.mu.Lock()
	defer m.mu.Unlock()

	var firstErr error
	for id, instance := range m.instances {
		if !instance.IsRunning {
			err := instance.Start()
			if err != nil {
				if firstErr == nil {
					firstErr = err
				}
			} else {
				log.Info().Msgf("INSTANCE MANAGER: Instancia %s iniciada.\n", id)
			}
		} else {
			log.Warn().Msgf("INSTANCE MANAGER: Instancia %s já está em execução.\n", id)
		}
	}
	return firstErr
}
