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

package models


type PanelProduct struct {

    /* 绑定tags，查数据必传 (Optional) */
    BindTags interface{} `json:"bindTags"`

    /* 子维度 (Optional) */
    Dimension []PanelDimension `json:"dimension"`

    /* 产品标识 (Optional) */
    Product string `json:"product"`

    /* 产品名称 (Optional) */
    ProductName string `json:"productName"`

    /* serviceCode (Optional) */
    ServiceCode string `json:"serviceCode"`

    /* 标签服务处注册的serviceCode (Optional) */
    TagServiceCode string `json:"tagServiceCode"`
}