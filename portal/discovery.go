package portal

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"distributed/pkg/config"
	"distributed/registry"
)

func BuildRegistration() registry.Registration {
	publicURL := config.GetEnv("PORTAL_PUBLIC_URL", "http://localhost:5000")

	return registry.Registration{
		ServiceName: registry.PortalService,
		ServiceURL:  publicURL,
		RequiredServices: []registry.ServiceName{
			registry.GradingService,
		},
		ServiceUpdateURL: publicURL + "/registry/updates",
		HeartBeatURL:     publicURL + "/healthz",
	}
}

func RegisterToRegistry(reg registry.Registration) error {
	registryBase := config.GetEnv("REGISTRY_URL", "http://localhost:3000")
	servicesURL := registryBase + "/services"

	buf := new(bytes.Buffer)
	if err := json.NewEncoder(buf).Encode(reg); err != nil {
		return err
	}

	res, err := http.Post(servicesURL, "application/json", buf)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("registry returned status %d on register", res.StatusCode)
	}

	return nil
}

func DeregisterFromRegistry(serviceURL string) error {
	registryBase := config.GetEnv("REGISTRY_URL", "http://localhost:3000")
	servicesURL := registryBase + "/services"

	req, err := http.NewRequest(http.MethodDelete, servicesURL, bytes.NewBuffer([]byte(serviceURL)))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "text/plain; charset=utf-8")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("registry returned status %d on deregister", res.StatusCode)
	}

	return nil
}
