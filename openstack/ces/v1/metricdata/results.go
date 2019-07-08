package metricdata

import (
	"github.com/gophercloud/gophercloud"
)

type MetricData struct {
	// Specifies the namespace in service.
	Namespace string `json:"namespace"`

	// The value can be a string of 1 to 64 characters
	// and must start with a letter and contain only uppercase
	// letters, lowercase letters, digits, and underscores.
	MetricName string `json:"metric_name"`

	//Specifies the list of the metric dimensions.
	Dimensions []map[string]interface{} `json:"dimensions"`

	// Specifies the metric data list.
	Datapoints []map[string]interface{} `json:"datapoints"`

	// Specifies the metric unit.
	Unit string `json:"unit"`
}

type MetricDatasResult struct {
	gophercloud.Result
}

// ExtractMetricDatas is a function that accepts a result and extracts metric datas.
func (r MetricDatasResult) ExtractMetricDatas() ([]MetricData, error) {
	var s struct {
		// Specifies the metric data.
		MetricDatas []MetricData `json:"metrics"`
	}
	err := r.ExtractInto(&s)
	return s.MetricDatas, err
}

type Datapoint struct {
	// ָ��ֵ�����ֶ����������������filterʹ�õĲ�ѯֵ��ͬ��
	Average float64 `json:"average"`
	// ָ��ɼ�ʱ�䡣
	Timestamp int `json:"timestamp"`
	// ָ�굥λ
	Unit string `json:"unit,omitempty"`
}

type EventDataInfo struct {
	// �¼����ͣ�����instance_host_info��
	Type string `json:"type"`
	// �¼��ϱ�ʱ�䡣
	Timestamp int `json:"timestamp"`
	// ����������Ϣ��
	Value string `json:"value"`
}

// This is a auto create Response Object
type EventData struct {
	Datapoints []EventDataInfo `json:"datapoints"`
}

type Metricdata struct {
	//  ָ�������б����ڲ�ѯ����ʱ���Ƽ�ػ������ѡ��ľۺ�������ǰȡ��from����������datapoints�а��������ݵ��п��ܻ����Ԥ�ڡ�
	Datapoints []Datapoint `json:"datapoints"`
	// ָ�����ƣ����絯���Ʒ��������ָ���е�cpu_util��
	MetricName string `json:"metric_name"`
}

type AddMetricDataResult struct {
	gophercloud.ErrResult
}

type GetEventDataResult struct {
	gophercloud.Result
}

type GetResult struct {
	gophercloud.Result
}

func (r GetEventDataResult) Extract() (*EventData, error) {
	var s *EventData
	err := r.ExtractInto(&s)
	return s, err
}

func (r GetResult) Extract() (*Metricdata, error) {
	var s *Metricdata
	err := r.ExtractInto(&s)
	return s, err
}
