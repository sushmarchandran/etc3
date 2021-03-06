/*
Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// metrics_builder.go - methods to programatically create metrics; used for testing

package v2alpha2

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// MetricBuilder type for building new config by hand
type MetricBuilder Metric

// NewMetric returns a new metric builder
func NewMetric(name, namespace string) *MetricBuilder {
	m := &Metric{
		TypeMeta: metav1.TypeMeta{
			APIVersion: GroupVersion.String(),
			Kind:       "Metric",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      name,
			Namespace: namespace,
		},
	}
	return (*MetricBuilder)(m)
}

// WithDescription ..
func (b *MetricBuilder) WithDescription(description string) *MetricBuilder {
	b.Spec.Description = &description
	return b
}

// WithParams ..
func (b *MetricBuilder) WithParams(params map[string]string) *MetricBuilder {
	paramsList := make([]NamedValue, 0)
	for name, value := range params {
		paramsList = append(paramsList, NamedValue{Name: name, Value: value})
	}
	b.Spec.Params = &paramsList
	return b
}

// WithUnits ..
func (b *MetricBuilder) WithUnits(units string) *MetricBuilder {
	b.Spec.Units = &units
	return b
}

// WithType ..
func (b *MetricBuilder) WithType(t MetricType) *MetricBuilder {
	b.Spec.Type = t
	return b
}

// WithProvider ..
func (b *MetricBuilder) WithProvider(provider string) *MetricBuilder {
	b.Spec.Provider = provider
	return b
}

// WithSampleSize ..
func (b *MetricBuilder) WithSampleSize(name string) *MetricBuilder {
	b.Spec.SampleSize = &name
	return b
}

// WithSecret ..
func (b *MetricBuilder) WithSecret(name string) *MetricBuilder {
	b.Spec.Secret = &name
	return b
}

// WithHeaders ..
func (b *MetricBuilder) WithHeaderTemplates(params map[string]string) *MetricBuilder {
	paramsList := make([]NamedValue, 0)
	for name, value := range params {
		paramsList = append(paramsList, NamedValue{Name: name, Value: value})
	}
	b.Spec.HeaderTemplates = &paramsList
	return b
}

// WithURLTemplate ..
func (b *MetricBuilder) WithURLTemplate(urlTemplate string) *MetricBuilder {
	b.Spec.URLTemplate = urlTemplate
	return b
}

// WithJQExpression ...
func (b *MetricBuilder) WithJQExpression(jqExpression string) *MetricBuilder {
	b.Spec.JQExpression = jqExpression
	return b
}

// Build ..
func (b *MetricBuilder) Build() *Metric {
	return (*Metric)(b)
}
