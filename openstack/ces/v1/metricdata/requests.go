package metricdata

import (
	"github.com/gophercloud/gophercloud"
)

// BatchQueryOptsBuilder allows extensions to add additional parameters to the
// BatchQuery request.
type BatchQueryOptsBuilder interface {
	ToBatchQueryOptsMap() (map[string]interface{}, error)
}

type Metric struct {
	// Specifies the namespace in service.
	Namespace string `json:"namespace" required:"true"`

	// The value can be a string of 1 to 64 characters
	// and must start with a letter and contain only uppercase
	// letters, lowercase letters, digits, and underscores.
	MetricName string `json:"metric_name" required:"true"`

	// Specifies the list of the metric dimensions.
	Dimensions []map[string]string `json:"dimensions" required:"true"`
}

// BatchQueryOpts represents options for batch query metric data.
type BatchQueryOpts struct {
	// Specifies the metric data.
	Metrics []Metric `json:"metrics" required:"true"`

	// Specifies the start time of the query.
	From int64 `json:"from" required:"true"`

	// Specifies the end time of the query.
	To int64 `json:"to" required:"true"`

	// Specifies the data monitoring granularity.
	Period string `json:"period" required:"true"`

	// Specifies the data rollup method.
	Filter string `json:"filter" required:"true"`
}

// ToBatchQueryOptsMap builds a request body from BatchQueryOpts.
func (opts BatchQueryOpts) ToBatchQueryOptsMap() (map[string]interface{}, error) {
	return gophercloud.BuildRequestBody(opts, "")
}

// Querying Monitoring Data in Batches.
func BatchQuery(client *gophercloud.ServiceClient, opts BatchQueryOptsBuilder) (r MetricDatasResult) {
	b, err := opts.ToBatchQueryOptsMap()
	if err != nil {
		r.Err = err
		return
	}
	_, r.Err = client.Post(batchQueryMetricDataURL(client), b, &r.Body, &gophercloud.RequestOpts{
		OkCodes: []int{200},})
	return
}

type AddMetricDataOpts []AddMetricDataItem

type GetEventDataOpts struct {
	// ָ���ά�ȣ�Ŀǰ���֧��3��ά�ȣ�ά�ȱ�Ŵ�0��ʼ��ά�ȸ�ʽΪdim.{i}=key,value�ο������Ʒ�����ά�ȡ�����dim.0=instance_id,i-12345
	Dim0 string `q:"dim.0,required"`
	Dim1 string `q:"dim.1"`
	Dim2 string `q:"dim.2"`
	// ��ѯ������ʼʱ�䣬UNIXʱ�������λ���롣
	From int `q:"from,required"`
	// ָ�������ռ䣬���絯���Ʒ����������ռ䡣
	Namespace string `q:"namespace,required"`
	// ��ѯ���ݽ�ֹʱ��UNIXʱ�������λ���롣from����С��to��
	To int `q:"to,required"`
	// �¼����ͣ�ֻ������ĸ���»��ߡ��л��ߣ���ĸ��ͷ�����Ȳ�����64����instance_host_info��
	Type string `q:"type,required"`
}

type GetOpts struct {
	// ָ���ά�ȣ�Ŀǰ���֧��3��ά�ȣ�ά�ȱ�Ŵ�0��ʼ��ά�ȸ�ʽΪdim.{i}=key,value�����ֵΪ256��  ����dim.0=instance_id,i-12345
	Dim0 string `q:"dim.0,required"`
	Dim1 string `q:"dim.1"`
	Dim2 string `q:"dim.2"`
	// ���ݾۺϷ�ʽ��  ֧�ֵ�ֵΪmax, min, average, sum, variance��
	Filter string `q:"filter,required"`
	// ��ѯ������ʼʱ�䣬UNIXʱ�������λ���롣����from��ֵ����ڵ�ǰʱ����ǰƫ������1�����ڡ����ھۺ�����Ĺ����ǽ�һ���ۺ����ڷ�Χ�ڵ����ݵ�ۺϵ�������ʼ�߽��ϣ������from��to�ķ�Χ�����ھۺ������ڣ�����Ϊ�ۺ�δ��ɶ���ɲ�ѯ����Ϊ�գ����Խ���from��������ڵ�ǰʱ����ǰƫ������1�����ڡ���5���Ӿۺ�����Ϊ�������赱ǰʱ���Ϊ10:35��10:30~10:35֮���ԭʼ���ݻᱻ�ۺϵ�10:30������ϣ����Բ�ѯ5�������ݵ�ʱfrom����ӦΪ10:30��֮ǰ���Ƽ�ػ������ѡ��ľۺ�������ǰȡ��from������
	From int `q:"from,required"`
	// ָ�����ƣ����絯���Ʒ��������ָ���е�cpu_util��
	MetricName string `q:"metric_name,required"`
	// ָ�������ռ䡣
	Namespace string `q:"namespace,required"`
	// ����������ȡ�  ȡֵ��Χ��  1��ʵʱ���� 300��5�������� 1200��20�������� 3600��1Сʱ���� 14400��4Сʱ���� 86400��1������
	Period int `q:"period,required"`
	// ��ѯ���ݽ�ֹʱ��UNIXʱ�������λ���롣from����С��to��
	To int `q:"to,required"`
}

type AddMetricDataItem struct {
	// ָ�����ݡ�
	Metric MetricInfo `json:"metric" required:"true"`
	// ���ݵ���Ч�ڣ���������Ч�����Զ�ɾ�������ݣ���λ�룬���ֵ604800��
	Ttl int `json:"ttl" required:"true"`
	// �����ռ�ʱ��  UNIXʱ�������λ���롣  ˵���� ��Ϊ�ͻ��˵�������������ʱ����˲������ݵ�ʱ���Ӧ����[��ǰʱ��-3��+20�룬��ǰʱ��+10����-20��]�����ڣ���֤���������ʱ������Ϊ����ʱ��������ݲ��ܲ������ݿ⡣
	CollectTime int `json:"collect_time" required:"true"`
	// ָ�����ݵ�ֵ��
	Value float64 `json:"value" required:"true"`
	// ���ݵĵ�λ��
	Unit string `json:"unit,omitempty"`
	// ���ݵ����ͣ�ֻ����\"int\"��\"float\"
	Type string `json:"type,omitempty"`
}

// ָ����Ϣ
type MetricInfo struct {
	// ָ��ά��
	Dimensions []MetricsDimension `json:"dimensions" required:"true"`
	// ָ�����ƣ���������ĸ��ͷ��ֻ�ܰ���0-9/a-z/A-Z/_���������Ϊ1�����Ϊ64��  ����ָ������μ���ѯָ���б��в�ѯ����ָ������
	MetricName string `json:"metric_name" required:"true"`
	// ָ�������ռ䣬�����絯���Ʒ����������ռ䡣��ʽΪservice.item��service��item�������ַ�������������ĸ��ͷ��ֻ�ܰ���0-9/a-z/A-Z/_���ܳ������Ϊ3�����Ϊ32��˵���� ��alarm_typeΪ��EVENT.SYS| EVENT.CUSTOM��ʱ����Ϊ�ա�
	Namespace string `json:"namespace" required:"true"`
	// ָ�굥λ
	Unit string `json:"unit,omitempty"`
}

// ָ��ά��
type MetricsDimension struct {
	// ά����
	Name string `json:"name,omitempty"`
	// ά��ֵ
	Value string `json:"value,omitempty"`
}

func (opts AddMetricDataItem) ToMap() (map[string]interface{}, error) {
	return gophercloud.BuildRequestBody(opts, "")
}

type AddMetricDataOptsBuilder interface {
	ToAddMetricDataMap() ([]map[string]interface{}, error)
}

func (opts AddMetricDataOpts) ToAddMetricDataMap() ([]map[string]interface{}, error) {
	newOpts := make([]map[string]interface{}, len(opts))
	for i, opt := range opts {
		opt, err := opt.ToMap()
		if err != nil {
			return nil, err
		}
		newOpts[i] = opt
	}
	return newOpts, nil
}

/*
func AddMetricData(client *gophercloud.ServiceClient, opts AddMetricDataOptsBuilder) (r AddMetricDataResult) {
	b, err := opts.ToAddMetricDataMap()
	if err != nil {
		r.Err = err
		return
	}

	_, r.Err = client.Post(addMetricDataURL(client), b, &r.Body, &gophercloud.RequestOpts{
		OkCodes: []int{201},
	})

	return
}


func GetEventData(client *gophercloud.ServiceClient, opts GetEventDataOpts) (r GetEventDataResult) {
	q, err := gophercloud.BuildQueryString(&opts)
	if err != nil {
		r.Err = err
		return
	}
	url := getEventDataURL(client) + q.String()
	_, r.Err = client.Get(url, &r.Body, &gophercloud.RequestOpts{
		OkCodes: []int{200},
	})

	return
}

func Get(client *gophercloud.ServiceClient, opts GetOpts) (r GetResult) {
	q, err := gophercloud.BuildQueryString(&opts)
	if err != nil {
		r.Err = err
		return
	}
	url := getURL(client) + q.String()
	_, r.Err = client.Get(url, &r.Body, &gophercloud.RequestOpts{
		OkCodes: []int{200},
	})

	return
}
*/