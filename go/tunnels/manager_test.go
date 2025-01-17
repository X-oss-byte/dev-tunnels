// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.

package tunnels

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/url"
	"os"
	"testing"
	"time"
)

var (
	serviceUrl           = ServiceProperties.ServiceURI
	ctx                  = context.Background()
	userAgentManagerTest = []UserAgent{{Name: "Tunnels-Go-SDK-Tests/Manager", Version: PackageVersion}}
)

func getUserToken() string {
	// Example: "github <gh-token>" or "Bearer <aad-token>"
	return ""
}

// These tests do not automatically run in the PR check github action
// beacuse they require authentication. If you want to run these tests
// you must first generate a tunnels access token and paste it in the
// getUserToken return value.
func TestTunnelCreateDelete(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping test in short mode.")
	}
	logger := log.New(os.Stdout, "", log.LstdFlags)

	url, err := url.Parse(serviceUrl)
	if err != nil {
		t.Errorf(err.Error())
	}

	managementClient, err := NewManager(userAgentManagerTest, getUserToken, url, nil)
	if err != nil {
		t.Errorf(err.Error())
	}

	tunnel := &Tunnel{}
	options := &TunnelRequestOptions{}
	createdTunnel, err := managementClient.CreateTunnel(ctx, tunnel, options)
	if err != nil {
		t.Errorf(err.Error())
		return
	}
	if createdTunnel.TunnelID == "" {
		t.Errorf("tunnel was not successfully created")
	} else {
		logger.Printf("Created tunnel with id %s", createdTunnel.TunnelID)
		createdTunnel.Table().Print()
	}

	err = managementClient.DeleteTunnel(ctx, createdTunnel, options)

	if err != nil {
		t.Errorf("tunnel was not successfully deleted")
	} else {
		logger.Printf("Deleted tunnel with id %s", createdTunnel.TunnelID)
	}
}

func TestListTunnels(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping test in short mode.")
	}
	logger := log.New(os.Stdout, "", log.LstdFlags)

	url, err := url.Parse(serviceUrl)
	if err != nil {
		t.Errorf(err.Error())
	}

	managementClient, err := NewManager(userAgentManagerTest, getUserToken, url, nil)
	if err != nil {
		t.Errorf(err.Error())
	}

	tunnel := &Tunnel{}
	options := &TunnelRequestOptions{}
	createdTunnel, err := managementClient.CreateTunnel(ctx, tunnel, options)
	if err != nil {
		t.Errorf(err.Error())
		return
	}
	if createdTunnel.TunnelID == "" {
		t.Errorf("tunnel was not successfully created")
	} else {
		logger.Printf("Created tunnel with id %s", createdTunnel.TunnelID)
		createdTunnel.Table().Print()
	}
	var token string
	if createdTunnel.AccessTokens != nil {
		token = createdTunnel.AccessTokens["manage"]
	} else {
		logger.Printf("Did not get token for created tunnel")
	}
	options = &TunnelRequestOptions{
		AccessToken: token,
	}
	tunnels, err := managementClient.ListTunnels(ctx, "", "", options)
	if err != nil {
		t.Errorf(err.Error())
	}
	if len(tunnels) == 0 {
		t.Errorf("tunnel was not successfully listed")
	}
	for _, tunnel := range tunnels {
		logger.Printf("found tunnel with id %s", tunnel.TunnelID)
		tunnel.Table().Print()
	}

	err = managementClient.DeleteTunnel(ctx, createdTunnel, options)

	if err != nil {
		t.Errorf("tunnel was not successfully deleted")
	} else {
		logger.Printf("Deleted tunnel with id %s", createdTunnel.TunnelID)
	}
}

func TestTunnelCreateUpdateDelete(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping test in short mode.")
	}
	logger := log.New(os.Stdout, "", log.LstdFlags)

	url, err := url.Parse(serviceUrl)
	if err != nil {
		t.Errorf(err.Error())
	}

	managementClient, err := NewManager(userAgentManagerTest, getUserToken, url, nil)
	if err != nil {
		t.Errorf(err.Error())
	}

	tunnel := &Tunnel{}
	options := &TunnelRequestOptions{}
	createdTunnel, err := managementClient.CreateTunnel(ctx, tunnel, options)
	if err != nil {
		t.Errorf(err.Error())
		return
	}
	if createdTunnel.TunnelID == "" {
		t.Errorf("tunnel was not successfully created")
	} else {
		logger.Printf("Created tunnel with id %s", createdTunnel.TunnelID)
		createdTunnel.Table().Print()
	}
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	generatedName := fmt.Sprintf("test%d", r1.Intn(10000))
	createdTunnel.Name = generatedName
	updatedTunnel, err := managementClient.UpdateTunnel(ctx, createdTunnel, []string{"Name"}, options)
	if err != nil {
		t.Errorf("tunnel was not successfully updated: %s", err.Error())
	} else if updatedTunnel.Name != generatedName {
		t.Errorf("tunnel was not successfully updated")
	} else {
		logger.Printf("Tunnel updated")
		updatedTunnel.Table().Print()
	}
	err = managementClient.DeleteTunnel(ctx, createdTunnel, options)

	if err != nil {
		t.Errorf("tunnel was not successfully deleted")
	} else {
		logger.Printf("Deleted tunnel with id %s", createdTunnel.TunnelID)
	}
}

func TestTunnelCreateUpdateTwiceDelete(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping test in short mode.")
	}
	logger := log.New(os.Stdout, "", log.LstdFlags)

	url, err := url.Parse(serviceUrl)
	if err != nil {
		t.Errorf(err.Error())
	}

	managementClient, err := NewManager(userAgentManagerTest, getUserToken, url, nil)
	if err != nil {
		t.Errorf(err.Error())
	}

	tunnel := &Tunnel{}
	options := &TunnelRequestOptions{}
	createdTunnel, err := managementClient.CreateTunnel(ctx, tunnel, options)
	if err != nil {
		t.Errorf(err.Error())
		return
	}
	if createdTunnel.TunnelID == "" {
		t.Errorf("tunnel was not successfully created")
	} else {
		logger.Printf("Created tunnel with id %s", createdTunnel.TunnelID)
		createdTunnel.Table().Print()
	}
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	generatedName := fmt.Sprintf("test%d", r1.Intn(10000))
	createdTunnel.Name = generatedName
	updatedTunnel, err := managementClient.UpdateTunnel(ctx, createdTunnel, []string{"Name"}, options)
	if err != nil {
		t.Errorf("tunnel was not successfully updated: %s", err.Error())
	} else if updatedTunnel.Name != generatedName {
		t.Errorf("tunnel was not successfully updated")
	} else {
		logger.Printf("Tunnel updated")
		updatedTunnel.Table().Print()
	}

	// In the second update we want to update the description without updating the name
	createdTunnel.Name = ""
	createdTunnel.Description = "test description"
	updatedTunnel, err = managementClient.UpdateTunnel(ctx, createdTunnel, []string{"Name", "Description"}, options)
	if err != nil {
		t.Errorf("tunnel was not successfully updated: %s", err.Error())
	} else if updatedTunnel.Name != generatedName || createdTunnel.Description != "test description" {
		t.Errorf("tunnel was not successfully updated")
	} else {
		logger.Printf("Tunnel updated")
		updatedTunnel.Table().Print()
	}
	err = managementClient.DeleteTunnel(ctx, createdTunnel, options)

	if err != nil {
		t.Errorf("tunnel was not successfully deleted")
	} else {
		logger.Printf("Deleted tunnel with id %s", createdTunnel.TunnelID)
	}
}

func TestTunnelCreateGetDelete(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping test in short mode.")
	}
	logger := log.New(os.Stdout, "", log.LstdFlags)

	url, err := url.Parse(serviceUrl)
	if err != nil {
		t.Errorf(err.Error())
	}

	managementClient, err := NewManager(userAgentManagerTest, getUserToken, url, nil)
	if err != nil {
		t.Errorf(err.Error())
	}

	tunnel := &Tunnel{}
	options := &TunnelRequestOptions{}
	createdTunnel, err := managementClient.CreateTunnel(ctx, tunnel, options)
	if err != nil {
		t.Errorf(err.Error())
		return
	}
	if createdTunnel.TunnelID == "" {
		t.Errorf("tunnel was not successfully created")
	} else {
		logger.Printf("Created tunnel with id %s", createdTunnel.TunnelID)
		createdTunnel.Table().Print()
	}

	getTunnel, err := managementClient.GetTunnel(ctx, createdTunnel, options)
	if err != nil {
		t.Errorf(err.Error())
		return
	}
	if getTunnel.TunnelID == "" {
		t.Errorf("tunnel was not successfully found")
	} else {
		logger.Printf("Got tunnel with id %s", getTunnel.TunnelID)
	}

	err = managementClient.DeleteTunnel(ctx, createdTunnel, options)

	if err != nil {
		t.Errorf("tunnel was not successfully deleted")
	} else {
		logger.Printf("Deleted tunnel with id %s", getTunnel.TunnelID)
	}
}

func TestTunnelAddPort(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping test in short mode.")
	}

	logger := log.New(os.Stdout, "", log.LstdFlags)

	url, err := url.Parse(serviceUrl)
	if err != nil {
		t.Errorf(err.Error())
	}

	managementClient, err := NewManager(userAgentManagerTest, getUserToken, url, nil)
	if err != nil {
		t.Errorf(err.Error())
	}

	tunnel := &Tunnel{}
	options := &TunnelRequestOptions{IncludePorts: true}
	createdTunnel, err := managementClient.CreateTunnel(ctx, tunnel, options)
	if err != nil {
		t.Errorf(err.Error())
		return
	}
	if createdTunnel.TunnelID == "" {
		t.Errorf("tunnel was not successfully created")
	} else {
		logger.Printf("Created tunnel with id %s", createdTunnel.TunnelID)
		createdTunnel.Table().Print()
	}
	portToAdd := NewTunnelPort(3000, "", "", "auto")
	port, err := managementClient.CreateTunnelPort(ctx, createdTunnel, portToAdd, options)
	if err != nil {
		t.Errorf(err.Error())
		return
	}
	logger.Printf("Created port: %d", port.PortNumber)
	port.Table().Print()

	getTunnel, err := managementClient.GetTunnel(ctx, createdTunnel, options)
	if err != nil {
		t.Errorf(err.Error())
		return
	}
	if getTunnel.TunnelID == "" {
		t.Errorf("tunnel was not successfully found")
	} else {
		logger.Printf("Got tunnel with id %s", getTunnel.TunnelID)
		getTunnel.Table().Print()
	}

	if len(getTunnel.Ports) != 1 {
		t.Errorf("port was not successfully added to tunnel")
	}

	err = managementClient.DeleteTunnel(ctx, createdTunnel, options)

	if err != nil {
		t.Errorf("tunnel was not successfully deleted")
	} else {
		logger.Printf("Deleted tunnel with id %s", createdTunnel.TunnelID)
	}
}

func TestTunnelDeletePort(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping test in short mode.")
	}

	logger := log.New(os.Stdout, "", log.LstdFlags)

	url, err := url.Parse(serviceUrl)
	if err != nil {
		t.Errorf(err.Error())
	}

	managementClient, err := NewManager(userAgentManagerTest, getUserToken, url, nil)
	if err != nil {
		t.Errorf(err.Error())
	}

	tunnel := &Tunnel{}
	options := &TunnelRequestOptions{IncludePorts: true}
	createdTunnel, err := managementClient.CreateTunnel(ctx, tunnel, options)
	if err != nil {
		t.Errorf(err.Error())
		return
	}
	if createdTunnel.TunnelID == "" {
		t.Errorf("tunnel was not successfully created")
	} else {
		logger.Printf("Created tunnel with id %s", createdTunnel.TunnelID)
		createdTunnel.Table().Print()
	}
	portToAdd := NewTunnelPort(3000, "", "", "auto")
	port, err := managementClient.CreateTunnelPort(ctx, createdTunnel, portToAdd, options)
	if err != nil {
		t.Errorf(err.Error())
		return
	}
	logger.Printf("Created port: %d", port.PortNumber)
	port.Table().Print()

	getTunnel, err := managementClient.GetTunnel(ctx, createdTunnel, options)
	if err != nil {
		t.Errorf(err.Error())
		return
	}
	if getTunnel.TunnelID == "" {
		t.Errorf("tunnel was not successfully found")
	} else {
		logger.Printf("Got tunnel with id %s", getTunnel.TunnelID)
		getTunnel.Table().Print()
	}

	err = managementClient.DeleteTunnelPort(ctx, createdTunnel, 3000, options)
	if err != nil {
		t.Errorf(err.Error())
		return
	}
	logger.Printf("Deleted port: %d", port.PortNumber)

	getTunnel, err = managementClient.GetTunnel(ctx, createdTunnel, options)
	if err != nil {
		t.Errorf(err.Error())
		return
	}
	if getTunnel.TunnelID == "" {
		t.Errorf("tunnel was not successfully found")
	} else {
		logger.Printf("Got tunnel with id %s", getTunnel.TunnelID)
		getTunnel.Table().Print()
	}

	if len(getTunnel.Ports) != 0 {
		t.Errorf("port was not successfully deleted")
	}

	err = managementClient.DeleteTunnel(ctx, createdTunnel, options)

	if err != nil {
		t.Errorf("tunnel was not successfully deleted")
	} else {
		logger.Printf("Deleted tunnel with id %s", createdTunnel.TunnelID)
	}
}

func TestTunnelUpdatePort(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping test in short mode.")
	}

	logger := log.New(os.Stdout, "", log.LstdFlags)

	url, err := url.Parse(serviceUrl)
	if err != nil {
		t.Errorf(err.Error())
	}

	managementClient, err := NewManager(userAgentManagerTest, getUserToken, url, nil)
	if err != nil {
		t.Errorf(err.Error())
	}

	tunnel := &Tunnel{}
	options := &TunnelRequestOptions{IncludePorts: true, TokenScopes: []TunnelAccessScope{"manage"}}
	createdTunnel, err := managementClient.CreateTunnel(ctx, tunnel, options)
	if err != nil {
		t.Errorf(err.Error())
		return
	}
	if createdTunnel.TunnelID == "" {
		t.Errorf("tunnel was not successfully created")
	} else {
		logger.Printf("Created tunnel with id %s", createdTunnel.TunnelID)
		createdTunnel.Table().Print()
	}
	portToAdd := NewTunnelPort(3000, "", "", "auto")
	port, err := managementClient.CreateTunnelPort(ctx, createdTunnel, portToAdd, options)
	if err != nil {
		t.Errorf(err.Error())
		return
	}
	logger.Printf("Created port: %d", port.PortNumber)
	port.Table().Print()

	getTunnel, err := managementClient.GetTunnel(ctx, createdTunnel, options)
	if err != nil {
		t.Errorf(err.Error())
		return
	}
	if getTunnel.TunnelID == "" {
		t.Errorf("tunnel was not successfully found")
	} else {
		logger.Printf("Got tunnel with id %s", getTunnel.TunnelID)
		getTunnel.Table().Print()
	}
	accessEntry := TunnelAccessControlEntry{
		Type:     TunnelAccessControlEntryTypeAnonymous,
		Subjects: []string{},
		Scopes:   []string{string(TunnelAccessScopeManage)},
	}
	portToAdd.AccessControl = &TunnelAccessControl{
		Entries: make([]TunnelAccessControlEntry, 0),
	}
	portToAdd.AccessControl.Entries = append(port.AccessControl.Entries, accessEntry)

	port, err = managementClient.UpdateTunnelPort(ctx, createdTunnel, portToAdd, nil, options)
	if err != nil {
		t.Errorf("port was not successfully updated: %s", err)
	} else if len(port.AccessControl.Entries) != 1 {
		t.Errorf("port was not successfully updated")
	}

	getTunnel, err = managementClient.GetTunnel(ctx, createdTunnel, options)
	if err != nil {
		t.Errorf(err.Error())
		return
	}
	if getTunnel.TunnelID == "" {
		t.Errorf("tunnel was not successfully found")
	} else {
		logger.Printf("Got tunnel with id %s", getTunnel.TunnelID)
		getTunnel.Table().Print()
	}
	if len(getTunnel.Ports[0].AccessControl.Entries) != 1 {
		t.Errorf("tunnel port was not successfully updated, access control was not changed")
	}

	port, err = managementClient.GetTunnelPort(ctx, createdTunnel, 3000, options)
	if err != nil {
		t.Errorf("port get error %s", err.Error())
		return
	}
	if len(port.AccessControl.Entries) != 1 {
		t.Errorf("tunnel port was not successfully updated, access control was not changed")
	}

	err = managementClient.DeleteTunnel(ctx, createdTunnel, options)

	if err != nil {
		t.Errorf("tunnel was not successfully deleted")
	} else {
		logger.Printf("Deleted tunnel with id %s", createdTunnel.TunnelID)
	}
}

func TestTunnelListPorts(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping test in short mode.")
	}

	logger := log.New(os.Stdout, "", log.LstdFlags)

	url, err := url.Parse(serviceUrl)
	if err != nil {
		t.Errorf(err.Error())
	}

	managementClient, err := NewManager(userAgentManagerTest, getUserToken, url, nil)
	if err != nil {
		t.Errorf(err.Error())
	}

	tunnel := &Tunnel{}
	options := &TunnelRequestOptions{IncludePorts: true}
	createdTunnel, err := managementClient.CreateTunnel(ctx, tunnel, options)
	if err != nil {
		t.Errorf(err.Error())
		return
	}
	if createdTunnel.TunnelID == "" {
		t.Errorf("tunnel was not successfully created")
	} else {
		logger.Printf("Created tunnel with id %s", createdTunnel.TunnelID)
		createdTunnel.Table().Print()
	}
	portToAdd := NewTunnelPort(3000, "", "", "auto")
	port, err := managementClient.CreateTunnelPort(ctx, createdTunnel, portToAdd, options)
	if err != nil {
		t.Errorf(err.Error())
		return
	}

	logger.Printf("Created port: %d", port.PortNumber)
	port.Table().Print()

	portToAdd = NewTunnelPort(3001, "", "", "auto")
	port, err = managementClient.CreateTunnelPort(ctx, createdTunnel, portToAdd, options)
	if err != nil {
		t.Errorf(err.Error())
		return
	}
	logger.Printf("Created port: %d", port.PortNumber)
	port.Table().Print()

	ports, err := managementClient.ListTunnelPorts(ctx, createdTunnel, options)
	if err != nil {
		t.Errorf(err.Error())
		return
	}
	if len(ports) != 2 {
		t.Errorf("ports not successfully listed")
	}
	for _, port := range ports {
		logger.Printf("Port: %d", port.PortNumber)
		port.Table().Print()
	}

	getTunnel, err := managementClient.GetTunnel(ctx, createdTunnel, options)
	if err != nil {
		t.Errorf(err.Error())
		return
	}
	if getTunnel.TunnelID == "" {
		t.Errorf("tunnel was not successfully found")
	} else {
		logger.Printf("Got tunnel with id %s", getTunnel.TunnelID)
		getTunnel.Table().Print()
	}

	if len(getTunnel.Ports) != 2 {
		t.Errorf("port was not successfully added to tunnel")
	}

	err = managementClient.DeleteTunnel(ctx, createdTunnel, options)

	if err != nil {
		t.Errorf("tunnel was not successfully deleted")
	} else {
		logger.Printf("Deleted tunnel with id %s", createdTunnel.TunnelID)
	}
}

func TestTunnelEndpoints(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping test in short mode.")
	}

	logger := log.New(os.Stdout, "", log.LstdFlags)

	url, err := url.Parse(serviceUrl)
	if err != nil {
		t.Errorf(err.Error())
	}

	managementClient, err := NewManager(userAgentManagerTest, getUserToken, url, nil)
	if err != nil {
		t.Errorf(err.Error())
	}

	tunnel := &Tunnel{}
	options := &TunnelRequestOptions{
		TokenScopes: managePortsAccessTokenScopes,
	}
	createdTunnel, err := managementClient.CreateTunnel(ctx, tunnel, options)
	if err != nil {
		t.Errorf(err.Error())
		return
	}
	if createdTunnel.TunnelID == "" {
		t.Errorf("tunnel was not successfully created")
	} else {
		logger.Printf("Created tunnel with id %s", createdTunnel.TunnelID)
		createdTunnel.Table().Print()
	}

	// Create and add endpoint
	endpoint := &TunnelEndpoint{
		HostID:         "test",
		ConnectionMode: TunnelConnectionModeTunnelRelay,
	}

	updatedEndpoint, err := managementClient.UpdateTunnelEndpoint(ctx, createdTunnel, endpoint, nil, options)

	if err != nil {
		t.Errorf(err.Error())
		return
	}
	logger.Printf("updated endpoint %s", updatedEndpoint.HostID)

	getTunnel, err := managementClient.GetTunnel(ctx, createdTunnel, options)
	if err != nil {
		t.Errorf(err.Error())
		return
	}
	if getTunnel.TunnelID == "" {
		t.Errorf("tunnel was not successfully found")
	} else {
		logger.Printf("Got tunnel with id %s", getTunnel.TunnelID)
	}
	if len(getTunnel.Endpoints) != 1 {
		t.Errorf("endpoint was not successfully updated")
	}

	err = managementClient.DeleteTunnelEndpoints(ctx, createdTunnel, "test", TunnelConnectionModeTunnelRelay, options)
	if err != nil {
		t.Errorf(err.Error())
		return
	}

	getTunnel, err = managementClient.GetTunnel(ctx, createdTunnel, options)
	if err != nil {
		t.Errorf(err.Error())
		return
	}
	if getTunnel.TunnelID == "" {
		t.Errorf("tunnel was not successfully found")
	} else {
		logger.Printf("Got tunnel with id %s", getTunnel.TunnelID)
	}
	if len(getTunnel.Endpoints) != 0 {
		t.Errorf("endpoint was not successfully deleted")
	}

	err = managementClient.DeleteTunnel(ctx, createdTunnel, options)

	if err != nil {
		t.Errorf("tunnel was not successfully deleted")
	} else {
		logger.Printf("Deleted tunnel with id %s", getTunnel.TunnelID)
	}
}


func TestResourceStatusUnmarshal(t *testing.T) {
	var test1 = []byte("{ \"current\": 3, \"limit\": 10 }")
	var result1 ResourceStatus
	var err = json.Unmarshal(test1, &result1)
	if err != nil {
		t.Error(err)
	}

	if (result1.Limit == 0) {
		t.Errorf("Limit was not deserialized")
	}

	var result2 ResourceStatus
	var test2 = []byte("3")
	err = json.Unmarshal(test2, &result2)
	if err != nil {
		t.Error(err)
	}

	if result1.Current != result2.Current {
		t.Errorf("%d != %d", result1.Current, result2.Current)
	}
}

func TestValidTokenScopes(t *testing.T) {
	var validScopes = TunnelAccessScopes{ "host", "connect" }
	var invalidScopes = TunnelAccessScopes{ "invalid", "connect" }
	var multiScopes = TunnelAccessScopes{ "host connect", "manage" }

	if err := validScopes.valid(nil, false); err != nil {
		t.Error(err)
	}
	if err := invalidScopes.valid(nil, false); err == nil {
		t.Errorf("Invalid scopes should not be valid")
	}
	if err := multiScopes.valid(nil, true); err != nil {
		t.Error(err)
	}
	if err := multiScopes.valid(nil, false); err == nil {
		t.Errorf("Multiple scopes should not be valid without allowMultiple flag")
	}
}
