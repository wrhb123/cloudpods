// Copyright 2018 JDCLOUD.COM
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// NOTE: This class is auto generated by the jdcloud code generator program.

package apis

import (
    "github.com/jdcloud-api/jdcloud-sdk-go/core"
    monitor "github.com/jdcloud-api/jdcloud-sdk-go/services/monitor/models"
)

type DescribeTagValuesRequest struct {

    core.JDCloudRequest

    /* region  */
    RegionId string `json:"regionId"`

    /* region  */
    TagKey string `json:"tagKey"`

    /* serviceCode  */
    ServiceCode string `json:"serviceCode"`

    /* 资源标识  */
    ResourceId string `json:"resourceId"`

    /* metric (Optional) */
    Metric *string `json:"metric"`

    /* 查询时间范围的开始时间(如不传,则默认为1时前)， UTC时间，格式：2016-12-11T00:00:00+0800（注意在url中+,:要转译,如2019-10-23T10%3A33%3A31%2B0800）。对于非连续的时序数据，开始时间无法保证准确性 (Optional) */
    StartTime *string `json:"startTime"`

    /* 查询时间范围的结束时间(如不传,则默认为当前)， UTC时间，格式：2016-12-11T00:00:00+0800（注意在url中+,:要转译,如2019-10-23T10%3A33%3A31%2B0800）.对于非连续的时序数据，结束时间无法保证准确性 (Optional) */
    EndTime *string `json:"endTime"`

    /* 时间间隔：1h，6h，12h，1d，3d，7d，14d，固定时间间隔，timeInterval默认为1h，当前时间往 前1h (Optional) */
    TimeInterval *string `json:"timeInterval"`

    /* 根据tags进行筛选(传入*和不传的效果一致) (Optional) */
    Tags []monitor.TagFilter `json:"tags"`
}

/*
 * param regionId: region (Required)
 * param tagKey: region (Required)
 * param serviceCode: serviceCode (Required)
 * param resourceId: 资源标识 (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewDescribeTagValuesRequest(
    regionId string,
    tagKey string,
    serviceCode string,
    resourceId string,
) *DescribeTagValuesRequest {

	return &DescribeTagValuesRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/tagValues/{tagKey}",
			Method:  "GET",
			Header:  nil,
			Version: "v2",
		},
        RegionId: regionId,
        TagKey: tagKey,
        ServiceCode: serviceCode,
        ResourceId: resourceId,
	}
}

/*
 * param regionId: region (Required)
 * param tagKey: region (Required)
 * param serviceCode: serviceCode (Required)
 * param resourceId: 资源标识 (Required)
 * param metric: metric (Optional)
 * param startTime: 查询时间范围的开始时间(如不传,则默认为1时前)， UTC时间，格式：2016-12-11T00:00:00+0800（注意在url中+,:要转译,如2019-10-23T10%3A33%3A31%2B0800）。对于非连续的时序数据，开始时间无法保证准确性 (Optional)
 * param endTime: 查询时间范围的结束时间(如不传,则默认为当前)， UTC时间，格式：2016-12-11T00:00:00+0800（注意在url中+,:要转译,如2019-10-23T10%3A33%3A31%2B0800）.对于非连续的时序数据，结束时间无法保证准确性 (Optional)
 * param timeInterval: 时间间隔：1h，6h，12h，1d，3d，7d，14d，固定时间间隔，timeInterval默认为1h，当前时间往 前1h (Optional)
 * param tags: 根据tags进行筛选(传入*和不传的效果一致) (Optional)
 */
func NewDescribeTagValuesRequestWithAllParams(
    regionId string,
    tagKey string,
    serviceCode string,
    resourceId string,
    metric *string,
    startTime *string,
    endTime *string,
    timeInterval *string,
    tags []monitor.TagFilter,
) *DescribeTagValuesRequest {

    return &DescribeTagValuesRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/tagValues/{tagKey}",
            Method:  "GET",
            Header:  nil,
            Version: "v2",
        },
        RegionId: regionId,
        TagKey: tagKey,
        ServiceCode: serviceCode,
        ResourceId: resourceId,
        Metric: metric,
        StartTime: startTime,
        EndTime: endTime,
        TimeInterval: timeInterval,
        Tags: tags,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewDescribeTagValuesRequestWithoutParam() *DescribeTagValuesRequest {

    return &DescribeTagValuesRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/tagValues/{tagKey}",
            Method:  "GET",
            Header:  nil,
            Version: "v2",
        },
    }
}

/* param regionId: region(Required) */
func (r *DescribeTagValuesRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param tagKey: region(Required) */
func (r *DescribeTagValuesRequest) SetTagKey(tagKey string) {
    r.TagKey = tagKey
}

/* param serviceCode: serviceCode(Required) */
func (r *DescribeTagValuesRequest) SetServiceCode(serviceCode string) {
    r.ServiceCode = serviceCode
}

/* param resourceId: 资源标识(Required) */
func (r *DescribeTagValuesRequest) SetResourceId(resourceId string) {
    r.ResourceId = resourceId
}

/* param metric: metric(Optional) */
func (r *DescribeTagValuesRequest) SetMetric(metric string) {
    r.Metric = &metric
}

/* param startTime: 查询时间范围的开始时间(如不传,则默认为1时前)， UTC时间，格式：2016-12-11T00:00:00+0800（注意在url中+,:要转译,如2019-10-23T10%3A33%3A31%2B0800）。对于非连续的时序数据，开始时间无法保证准确性(Optional) */
func (r *DescribeTagValuesRequest) SetStartTime(startTime string) {
    r.StartTime = &startTime
}

/* param endTime: 查询时间范围的结束时间(如不传,则默认为当前)， UTC时间，格式：2016-12-11T00:00:00+0800（注意在url中+,:要转译,如2019-10-23T10%3A33%3A31%2B0800）.对于非连续的时序数据，结束时间无法保证准确性(Optional) */
func (r *DescribeTagValuesRequest) SetEndTime(endTime string) {
    r.EndTime = &endTime
}

/* param timeInterval: 时间间隔：1h，6h，12h，1d，3d，7d，14d，固定时间间隔，timeInterval默认为1h，当前时间往 前1h(Optional) */
func (r *DescribeTagValuesRequest) SetTimeInterval(timeInterval string) {
    r.TimeInterval = &timeInterval
}

/* param tags: 根据tags进行筛选(传入*和不传的效果一致)(Optional) */
func (r *DescribeTagValuesRequest) SetTags(tags []monitor.TagFilter) {
    r.Tags = tags
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r DescribeTagValuesRequest) GetRegionId() string {
    return r.RegionId
}

type DescribeTagValuesResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result DescribeTagValuesResult `json:"result"`
}

type DescribeTagValuesResult struct {
    TagValues []string `json:"tagValues"`
}