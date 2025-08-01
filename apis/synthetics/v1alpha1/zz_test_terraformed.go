/*
Copyright 2022 Upbound Inc.
*/

// Code generated by upjet. DO NOT EDIT.

package v1alpha1

import (
	"dario.cat/mergo"
	"github.com/pkg/errors"

	"github.com/crossplane/upjet/pkg/resource"
	"github.com/crossplane/upjet/pkg/resource/json"
)

// GetTerraformResourceType returns Terraform resource type for this Test
func (mg *Test) GetTerraformResourceType() string {
	return "datadog_synthetics_test"
}

// GetConnectionDetailsMapping for this Test
func (tr *Test) GetConnectionDetailsMapping() map[string]string {
	return map[string]string{"api_step[*].request_basicauth[*].access_key": "apiStep[*].requestBasicauth[*].accessKeySecretRef", "api_step[*].request_basicauth[*].client_secret": "apiStep[*].requestBasicauth[*].clientSecretSecretRef", "api_step[*].request_basicauth[*].password": "apiStep[*].requestBasicauth[*].passwordSecretRef", "api_step[*].request_basicauth[*].secret_key": "apiStep[*].requestBasicauth[*].secretKeySecretRef", "api_step[*].request_client_certificate[*].cert[*].content": "apiStep[*].requestClientCertificate[*].cert[*].contentSecretRef", "api_step[*].request_client_certificate[*].key[*].content": "apiStep[*].requestClientCertificate[*].key[*].contentSecretRef", "request_basicauth[*].access_key": "requestBasicauth[*].accessKeySecretRef", "request_basicauth[*].client_secret": "requestBasicauth[*].clientSecretSecretRef", "request_basicauth[*].password": "requestBasicauth[*].passwordSecretRef", "request_basicauth[*].secret_key": "requestBasicauth[*].secretKeySecretRef", "request_client_certificate[*].cert[*].content": "requestClientCertificate[*].cert[*].contentSecretRef", "request_client_certificate[*].key[*].content": "requestClientCertificate[*].key[*].contentSecretRef"}
}

// GetObservation of this Test
func (tr *Test) GetObservation() (map[string]any, error) {
	o, err := json.TFParser.Marshal(tr.Status.AtProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(o, &base)
}

// SetObservation for this Test
func (tr *Test) SetObservation(obs map[string]any) error {
	p, err := json.TFParser.Marshal(obs)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Status.AtProvider)
}

// GetID returns ID of underlying Terraform resource of this Test
func (tr *Test) GetID() string {
	if tr.Status.AtProvider.ID == nil {
		return ""
	}
	return *tr.Status.AtProvider.ID
}

// GetParameters of this Test
func (tr *Test) GetParameters() (map[string]any, error) {
	p, err := json.TFParser.Marshal(tr.Spec.ForProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(p, &base)
}

// SetParameters for this Test
func (tr *Test) SetParameters(params map[string]any) error {
	p, err := json.TFParser.Marshal(params)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Spec.ForProvider)
}

// GetInitParameters of this Test
func (tr *Test) GetInitParameters() (map[string]any, error) {
	p, err := json.TFParser.Marshal(tr.Spec.InitProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(p, &base)
}

// GetInitParameters of this Test
func (tr *Test) GetMergedParameters(shouldMergeInitProvider bool) (map[string]any, error) {
	params, err := tr.GetParameters()
	if err != nil {
		return nil, errors.Wrapf(err, "cannot get parameters for resource '%q'", tr.GetName())
	}
	if !shouldMergeInitProvider {
		return params, nil
	}

	initParams, err := tr.GetInitParameters()
	if err != nil {
		return nil, errors.Wrapf(err, "cannot get init parameters for resource '%q'", tr.GetName())
	}

	// Note(lsviben): mergo.WithSliceDeepCopy is needed to merge the
	// slices from the initProvider to forProvider. As it also sets
	// overwrite to true, we need to set it back to false, we don't
	// want to overwrite the forProvider fields with the initProvider
	// fields.
	err = mergo.Merge(&params, initParams, mergo.WithSliceDeepCopy, func(c *mergo.Config) {
		c.Overwrite = false
	})
	if err != nil {
		return nil, errors.Wrapf(err, "cannot merge spec.initProvider and spec.forProvider parameters for resource '%q'", tr.GetName())
	}

	return params, nil
}

// LateInitialize this Test using its observed tfState.
// returns True if there are any spec changes for the resource.
func (tr *Test) LateInitialize(attrs []byte) (bool, error) {
	params := &TestParameters{}
	if err := json.TFParser.Unmarshal(attrs, params); err != nil {
		return false, errors.Wrap(err, "failed to unmarshal Terraform state parameters for late-initialization")
	}
	opts := []resource.GenericLateInitializerOption{resource.WithZeroValueJSONOmitEmptyFilter(resource.CNameWildcard)}

	li := resource.NewGenericLateInitializer(opts...)
	return li.LateInitialize(&tr.Spec.ForProvider, params)
}

// GetTerraformSchemaVersion returns the associated Terraform schema version
func (tr *Test) GetTerraformSchemaVersion() int {
	return 0
}
