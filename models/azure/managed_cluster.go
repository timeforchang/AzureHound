// Copyright (C) 2022 Specter Ops, Inc.
//
// This file is part of AzureHound.
//
// AzureHound is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// AzureHound is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program.  If not, see <https://www.gnu.org/licenses/>.

package azure

import "strings"

type ManagedCluster struct {
	Entity

	ExtendedLocation ExtendedLocation              `json:"extendedLocation,omitempty"`
	Identity         ManagedClusterManagedIdentity `json:"identity,omitempty"`
	Properties       ManagedClusterProperties      `json:"properties,omitempty"`
	Location         string                        `json:"location,omitempty"`
	Name             string                        `json:"name,omitempty"`
	Plan             Plan                          `json:"plan,omitempty"`
	Tags             map[string]string             `json:"tags,omitempty"`
	Type             string                        `json:"type,omitempty"`
	Zones            []string                      `json:"zones,omitempty"`
}

func (s *ManagedCluster) PopulateManagedIdentity() {
	if s.Properties.ServicePrincipalProfile != (ManagedClusterServicePrincipalProfile{}) {
		s.Identity.ServicePrincipalProfile = s.Properties.ServicePrincipalProfile
		s.Identity.IdentityProfile = s.Properties.IdentityProfile
	}
}

func (s ManagedCluster) ResourceGroupName() string {
	parts := strings.Split(s.Id, "/")
	if len(parts) > 4 {
		return parts[4]
	} else {
		return ""
	}
}

func (s ManagedCluster) ResourceGroupId() string {
	parts := strings.Split(s.Id, "/")
	if len(parts) > 5 {
		return strings.Join(parts[:5], "/")
	} else {
		return ""
	}
}
