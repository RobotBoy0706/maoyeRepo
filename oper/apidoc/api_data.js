define({ "api": [
  {
    "success": {
      "fields": {
        "Success 200": [
          {
            "group": "Success 200",
            "optional": false,
            "field": "varname1",
            "description": "<p>No type.</p>"
          },
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "varname2",
            "description": "<p>With type.</p>"
          }
        ]
      }
    },
    "type": "",
    "url": "",
    "version": "0.0.0",
    "filename": "./apidoc/main.js",
    "group": "/home/zhaolunxiang/go/src/dg/maoyetrpg-back/oper/apidoc/main.js",
    "groupTitle": "/home/zhaolunxiang/go/src/dg/maoyetrpg-back/oper/apidoc/main.js",
    "name": ""
  },
  {
    "type": "",
    "url": "DELETE",
    "title": "/api/v1/advertisement/advertisement 删除广告",
    "group": "advertisement",
    "query": [
      {
        "group": "Query",
        "optional": false,
        "field": "userid",
        "description": "<p>int 必填，执行请求的用户id（只有管理员用户才能删除,管理员id在config.json中admin处设置）</p>"
      },
      {
        "group": "Query",
        "optional": false,
        "field": "token",
        "description": "<p>string 必填，执行请求的用户token</p>"
      },
      {
        "group": "Query",
        "optional": false,
        "field": "id",
        "description": "<p>string 必填，要删除的广告id { &quot;data&quot;: { &quot;id&quot;: 28, &quot;title&quot;: &quot;abcd&quot;, &quot;text&quot;: &quot;广告，公告管理系统&quot;, &quot;type&quot;: &quot;广告&quot;, &quot;picturename&quot;: &quot;接口信息.doc&quot;, &quot;md5&quot;: &quot;f77e06096252c8a27236f137e0f0dd89&quot;, &quot;starttime&quot;: &quot;2021-01-02 15:04:05&quot;, &quot;endtime&quot;: &quot;2022-01-02 15:04:05&quot; }, &quot;errcode&quot;: 0, &quot;errmsg&quot;: &quot;ok&quot; }</p>"
      }
    ],
    "error": {
      "fields": {
        "Error 4xx": [
          {
            "group": "Error 4xx",
            "optional": false,
            "field": "200",
            "description": "<p>OK</p>"
          }
        ]
      }
    },
    "parameter": {
      "fields": {
        "Parameter": [
          {
            "group": "Parameter",
            "optional": false,
            "field": "errcode",
            "description": "<p>int 错误代码</p>"
          },
          {
            "group": "Parameter",
            "optional": false,
            "field": "errmsg",
            "description": "<p>string 错误信息</p>"
          }
        ]
      }
    },
    "examples": [
      {
        "title": "json",
        "content": "服务器内部错误\n{\n    \"errcode\":30001,\n    \"errmsg\":\"system internal error\"\n}",
        "type": "json"
      },
      {
        "title": "json",
        "content": "用户不合法\n{\n    \"errcode\":30305,\n    \"errmsg\":\"auth failed\"\n}",
        "type": "json"
      },
      {
        "title": "json",
        "content": "参数错误\n{\n    \"errcode\":30002,\n    \"errmsg\":\"invalid parameter\"\n}\n不是管理员\n{\n    \"errcode\":30005,\n    \"errmsg\":\"no permission\"\n}",
        "type": "json"
      },
      {
        "title": "json",
        "content": "操作数据库失败\n{\n    \"errcode\":30009,\n    \"errmsg\":\"database operation error\"\n}",
        "type": "json"
      }
    ],
    "version": "0.0.0",
    "filename": "./advertisement.go",
    "groupTitle": "advertisement",
    "name": "Delete"
  },
  {
    "type": "",
    "url": "GET",
    "title": "/api/v1/advertisement/advertisementPicture 获取广告图片",
    "group": "advertisement",
    "query": [
      {
        "group": "Query",
        "optional": false,
        "field": "userid",
        "description": "<p>int 必填，执行请求的用户id（这里可以用非管理员id，不过非管理员只能看到未过期的广告，管理员可以看到全部广告,管理员id在config.json中admin处设置）</p>"
      },
      {
        "group": "Query",
        "optional": false,
        "field": "token",
        "description": "<p>string 必填，执行请求的用户token</p>"
      },
      {
        "group": "Query",
        "optional": false,
        "field": "md5",
        "description": "<p>string 必填，图片的md5值，这个可以通过上面的获取广告信息中md5得到 返回的是图片</p>"
      }
    ],
    "error": {
      "fields": {
        "Error 4xx": [
          {
            "group": "Error 4xx",
            "optional": false,
            "field": "200",
            "description": "<p>OK</p>"
          }
        ]
      }
    },
    "parameter": {
      "fields": {
        "Parameter": [
          {
            "group": "Parameter",
            "optional": false,
            "field": "errcode",
            "description": "<p>int 错误代码</p>"
          },
          {
            "group": "Parameter",
            "optional": false,
            "field": "errmsg",
            "description": "<p>string 错误信息</p>"
          }
        ]
      }
    },
    "examples": [
      {
        "title": "json",
        "content": "服务器内部错误\n{\n    \"errcode\":30001,\n    \"errmsg\":\"system internal error\"\n}",
        "type": "json"
      },
      {
        "title": "json",
        "content": "用户不合法\n{\n    \"errcode\":30305,\n    \"errmsg\":\"auth failed\"\n}",
        "type": "json"
      },
      {
        "title": "json",
        "content": "参数错误\n{\n    \"errcode\":30002,\n    \"errmsg\":\"invalid parameter\"\n}",
        "type": "json"
      },
      {
        "title": "json",
        "content": "操作数据库失败\n{\n    \"errcode\":30009,\n    \"errmsg\":\"database operation error\"\n}",
        "type": "json"
      }
    ],
    "version": "0.0.0",
    "filename": "./advertisement.go",
    "groupTitle": "advertisement",
    "name": "Get"
  },
  {
    "type": "",
    "url": "GET",
    "title": "/api/v1/advertisement/advertisement 获取广告",
    "group": "advertisement",
    "query": [
      {
        "group": "Query",
        "optional": false,
        "field": "userid",
        "description": "<p>int 必填，执行请求的用户id（这里可以用非管理员id，不过非管理员只能看到未过期的广告，管理员可以看到全部广告,管理员id在config.json中admin处设置）</p>"
      },
      {
        "group": "Query",
        "optional": false,
        "field": "token",
        "description": "<p>string 必填，执行请求的用户token 返回的是请求有权限看到的广告列表 { &quot;data&quot;: { &quot;id&quot;: 28, &quot;title&quot;: &quot;abcd&quot;, &quot;text&quot;: &quot;广告，公告管理系统&quot;, &quot;type&quot;: &quot;广告&quot;, &quot;picturename&quot;: &quot;接口信息.doc&quot;, &quot;md5&quot;: &quot;f77e06096252c8a27236f137e0f0dd89&quot;, &quot;starttime&quot;: &quot;2021-01-02 15:04:05&quot;, &quot;endtime&quot;: &quot;2022-01-02 15:04:05&quot; }, &quot;errcode&quot;: 0, &quot;errmsg&quot;: &quot;ok&quot; } { &quot;data&quot;: { &quot;id&quot;: 30, &quot;title&quot;: &quot;abcd&quot;, &quot;text&quot;: &quot;广告，公告管理系统&quot;, &quot;type&quot;: &quot;广告&quot;, &quot;picturename&quot;: &quot;接口信息.doc&quot;, &quot;md5&quot;: &quot;f77e06096252c8a27236f137e0f0dd89&quot;, &quot;starttime&quot;: &quot;2021-01-02 15:04:05&quot;, &quot;endtime&quot;: &quot;2022-01-02 15:04:05&quot; }, &quot;errcode&quot;: 0, &quot;errmsg&quot;: &quot;ok&quot; }</p>"
      }
    ],
    "error": {
      "fields": {
        "Error 4xx": [
          {
            "group": "Error 4xx",
            "optional": false,
            "field": "200",
            "description": "<p>OK</p>"
          }
        ]
      }
    },
    "parameter": {
      "fields": {
        "Parameter": [
          {
            "group": "Parameter",
            "optional": false,
            "field": "errcode",
            "description": "<p>int 错误代码</p>"
          },
          {
            "group": "Parameter",
            "optional": false,
            "field": "errmsg",
            "description": "<p>string 错误信息</p>"
          }
        ]
      }
    },
    "examples": [
      {
        "title": "json",
        "content": "服务器内部错误\n{\n    \"errcode\":30001,\n    \"errmsg\":\"system internal error\"\n}",
        "type": "json"
      },
      {
        "title": "json",
        "content": "用户不合法\n{\n    \"errcode\":30305,\n    \"errmsg\":\"auth failed\"\n}",
        "type": "json"
      },
      {
        "title": "json",
        "content": "参数错误\n{\n    \"errcode\":30002,\n    \"errmsg\":\"invalid parameter\"\n}",
        "type": "json"
      },
      {
        "title": "json",
        "content": "操作数据库失败\n{\n    \"errcode\":30009,\n    \"errmsg\":\"database operation error\"\n}",
        "type": "json"
      }
    ],
    "version": "0.0.0",
    "filename": "./advertisement.go",
    "groupTitle": "advertisement",
    "name": "Get"
  },
  {
    "type": "",
    "url": "POST",
    "title": "/api/v1/advertisement/advertisement 新增广告",
    "group": "advertisement",
    "query": [
      {
        "group": "Query",
        "optional": false,
        "field": "userid",
        "description": "<p>int 必填，执行请求的用户id（只有管理员用户才能新增,管理员id在config.json中admin处设置）</p>"
      },
      {
        "group": "Query",
        "optional": false,
        "field": "token",
        "description": "<p>string 必填，执行请求的用户token</p>"
      }
    ],
    "parameter": {
      "fields": {
        "Parameter": [
          {
            "group": "Parameter",
            "optional": false,
            "field": "title",
            "description": "<p>string 必填，标题</p>"
          },
          {
            "group": "Parameter",
            "optional": false,
            "field": "text",
            "description": "<p>string 必填，文本</p>"
          },
          {
            "group": "Parameter",
            "optional": false,
            "field": "md5",
            "description": "<p>string 必填，资源md5</p>"
          },
          {
            "group": "Parameter",
            "optional": false,
            "field": "file",
            "description": "<p>file 必填，文件数据流</p>"
          },
          {
            "group": "Parameter",
            "optional": false,
            "field": "starttime",
            "description": "<p>string 必填 开始生效时间 必须用2021-01-02 15:04:05的格式</p>"
          },
          {
            "group": "Parameter",
            "optional": false,
            "field": "endtime",
            "description": "<p>string 必填 结束生效时间 必须用2021-01-02 15:04:05的格式</p>"
          },
          {
            "group": "Parameter",
            "optional": false,
            "field": "type",
            "description": "<p>string 必填 类型 广告/公告 { &quot;data&quot;: { &quot;id&quot;: 28, &quot;title&quot;: &quot;abcd&quot;, &quot;text&quot;: &quot;广告，公告管理系统&quot;, &quot;type&quot;: &quot;广告&quot;, &quot;picturename&quot;: &quot;接口信息.doc&quot;, &quot;md5&quot;: &quot;f77e06096252c8a27236f137e0f0dd89&quot;, &quot;starttime&quot;: &quot;2021-01-02 15:04:05&quot;, &quot;endtime&quot;: &quot;2022-01-02 15:04:05&quot; }, &quot;errcode&quot;: 0, &quot;errmsg&quot;: &quot;ok&quot; }</p>"
          },
          {
            "group": "Parameter",
            "optional": false,
            "field": "errcode",
            "description": "<p>int 错误代码</p>"
          },
          {
            "group": "Parameter",
            "optional": false,
            "field": "errmsg",
            "description": "<p>string 错误信息</p>"
          }
        ]
      }
    },
    "error": {
      "fields": {
        "Error 4xx": [
          {
            "group": "Error 4xx",
            "optional": false,
            "field": "200",
            "description": "<p>OK</p>"
          }
        ]
      }
    },
    "examples": [
      {
        "title": "json",
        "content": "服务器内部错误\n{\n    \"errcode\":30001,\n    \"errmsg\":\"system internal error\"\n}",
        "type": "json"
      },
      {
        "title": "json",
        "content": "用户不合法\n{\n    \"errcode\":30305,\n    \"errmsg\":\"auth failed\"\n}",
        "type": "json"
      },
      {
        "title": "json",
        "content": "参数错误\n{\n    \"errcode\":30002,\n    \"errmsg\":\"invalid parameter\"\n}\n不是管理员\n{\n    \"errcode\":30005,\n    \"errmsg\":\"no permission\"\n}",
        "type": "json"
      },
      {
        "title": "json",
        "content": "操作数据库失败\n{\n    \"errcode\":30009,\n    \"errmsg\":\"database operation error\"\n}",
        "type": "json"
      }
    ],
    "version": "0.0.0",
    "filename": "./advertisement.go",
    "groupTitle": "advertisement",
    "name": "Post"
  },
  {
    "type": "",
    "url": "PUT",
    "title": "/api/v1/advertisement/advertisement 新增广告",
    "group": "advertisement",
    "query": [
      {
        "group": "Query",
        "optional": false,
        "field": "userid",
        "description": "<p>int 必填，执行请求的用户id（只有管理员用户才能新增,管理员id在config.json中admin处设置）</p>"
      },
      {
        "group": "Query",
        "optional": false,
        "field": "token",
        "description": "<p>string 必填，执行请求的用户token</p>"
      }
    ],
    "parameter": {
      "fields": {
        "Parameter": [
          {
            "group": "Parameter",
            "optional": false,
            "field": "title",
            "description": "<p>string 必填，标题</p>"
          },
          {
            "group": "Parameter",
            "optional": false,
            "field": "text",
            "description": "<p>string 必填，文本</p>"
          },
          {
            "group": "Parameter",
            "optional": false,
            "field": "md5",
            "description": "<p>string 必填，资源md5</p>"
          },
          {
            "group": "Parameter",
            "optional": false,
            "field": "file",
            "description": "<p>file 必填，文件数据流</p>"
          },
          {
            "group": "Parameter",
            "optional": false,
            "field": "starttime",
            "description": "<p>string 必填 开始生效时间 必须用2021-01-02 15:04:05的格式</p>"
          },
          {
            "group": "Parameter",
            "optional": false,
            "field": "endtime",
            "description": "<p>string 必填 结束生效时间 必须用2021-01-02 15:04:05的格式</p>"
          },
          {
            "group": "Parameter",
            "optional": false,
            "field": "type",
            "description": "<p>string 必填 类型 广告/公告 { &quot;data&quot;: { &quot;id&quot;: 28, &quot;title&quot;: &quot;abcd&quot;, &quot;text&quot;: &quot;广告，公告管理系统&quot;, &quot;type&quot;: &quot;广告&quot;, &quot;picturename&quot;: &quot;接口信息.doc&quot;, &quot;md5&quot;: &quot;f77e06096252c8a27236f137e0f0dd89&quot;, &quot;starttime&quot;: &quot;2021-01-02 15:04:05&quot;, &quot;endtime&quot;: &quot;2022-01-02 15:04:05&quot; }, &quot;errcode&quot;: 0, &quot;errmsg&quot;: &quot;ok&quot; }</p>"
          },
          {
            "group": "Parameter",
            "optional": false,
            "field": "errcode",
            "description": "<p>int 错误代码</p>"
          },
          {
            "group": "Parameter",
            "optional": false,
            "field": "errmsg",
            "description": "<p>string 错误信息</p>"
          }
        ]
      }
    },
    "error": {
      "fields": {
        "Error 4xx": [
          {
            "group": "Error 4xx",
            "optional": false,
            "field": "200",
            "description": "<p>OK</p>"
          }
        ]
      }
    },
    "examples": [
      {
        "title": "json",
        "content": "服务器内部错误\n{\n    \"errcode\":30001,\n    \"errmsg\":\"system internal error\"\n}",
        "type": "json"
      },
      {
        "title": "json",
        "content": "用户不合法\n{\n    \"errcode\":30305,\n    \"errmsg\":\"auth failed\"\n}",
        "type": "json"
      },
      {
        "title": "json",
        "content": "参数错误\n{\n    \"errcode\":30002,\n    \"errmsg\":\"invalid parameter\"\n}\n不是管理员\n{\n    \"errcode\":30005,\n    \"errmsg\":\"no permission\"\n}",
        "type": "json"
      },
      {
        "title": "json",
        "content": "操作数据库失败\n{\n    \"errcode\":30009,\n    \"errmsg\":\"database operation error\"\n}",
        "type": "json"
      }
    ],
    "version": "0.0.0",
    "filename": "./advertisement.go",
    "groupTitle": "advertisement",
    "name": "Put"
  },
  {
    "type": "",
    "url": "DELETE /api/v1/resource/resource 删除一个资源",
    "title": "只有管理员能删除类型type为站内新闻的资源",
    "group": "resource",
    "query": [
      {
        "group": "Query",
        "optional": false,
        "field": "userid",
        "description": "<p>int 必填，执行请求的用户id</p>"
      },
      {
        "group": "Query",
        "optional": false,
        "field": "token",
        "description": "<p>string 必填，执行请求的用户token</p>"
      },
      {
        "group": "Query",
        "optional": false,
        "field": "id",
        "description": "<p>int64 必填，需要删除的资源的id</p>"
      },
      {
        "group": "Query",
        "optional": false,
        "field": "layer_id",
        "description": "<p>int64 必填，需要删除的layer id</p>"
      },
      {
        "group": "Query",
        "optional": false,
        "field": "state",
        "description": "<p>int64 必填，槽id</p>"
      }
    ],
    "success": {
      "fields": {
        "Success 200": [
          {
            "group": "Success 200",
            "optional": false,
            "field": "200",
            "description": "<p>OK</p>"
          }
        ]
      }
    },
    "parameter": {
      "fields": {
        "Parameter": [
          {
            "group": "Parameter",
            "optional": false,
            "field": "errcode",
            "description": "<p>int 错误代码</p>"
          },
          {
            "group": "Parameter",
            "optional": false,
            "field": "errmsg",
            "description": "<p>string 错误信息</p>"
          }
        ]
      }
    },
    "examples": [
      {
        "title": "json",
        "content": "{\n    \"errcode\":0,\n    \"errmsg\":\"ok\"\n}",
        "type": "json"
      },
      {
        "title": "json",
        "content": "服务器内部错误\n{\n    \"errcode\":30001,\n    \"errmsg\":\"system internal error\"\n}",
        "type": "json"
      },
      {
        "title": "json",
        "content": "用户不合法\n{\n    \"errcode\":30305,\n    \"errmsg\":\"auth failed\"\n}",
        "type": "json"
      },
      {
        "title": "json",
        "content": "参数错误\n{\n    \"errcode\":30002,\n    \"errmsg\":\"invalid parameter\"\n}",
        "type": "json"
      },
      {
        "title": "json",
        "content": "操作数据库失败\n{\n    \"errcode\":30009,\n    \"errmsg\":\"database operation error\"\n}",
        "type": "json"
      }
    ],
    "error": {
      "fields": {
        "Error 4xx": [
          {
            "group": "Error 4xx",
            "optional": false,
            "field": "200",
            "description": "<p>OK</p>"
          }
        ]
      }
    },
    "version": "0.0.0",
    "filename": "./resource.go",
    "groupTitle": "resource",
    "name": "DeleteApiV1ResourceResource"
  },
  {
    "type": "",
    "url": "GET",
    "title": "/api/v1/resource/data GET方法获取资源(支持Content-Range，如果resource的type是音乐的话，Content-Type会设置成audio/mpeg)",
    "group": "resource",
    "query": [
      {
        "group": "Query",
        "optional": false,
        "field": "userid",
        "description": "<p>int 必填，执行请求的用户id</p>"
      },
      {
        "group": "Query",
        "optional": false,
        "field": "token",
        "description": "<p>string 必填，执行请求的用户token</p>"
      }
    ],
    "parameter": {
      "fields": {
        "Parameter": [
          {
            "group": "Parameter",
            "optional": false,
            "field": "md5",
            "description": "<p>string 必填，资源md5值</p>"
          },
          {
            "group": "Parameter",
            "optional": false,
            "field": "timestamp",
            "description": "<p>int 选填，请求时时间戳</p>"
          },
          {
            "group": "Parameter",
            "optional": false,
            "field": "download",
            "description": "<p>int 选填，是否下载，1为下载，0为不下载，默认为0</p>"
          },
          {
            "group": "Parameter",
            "optional": false,
            "field": "file_name",
            "description": "<p>string 选填，文件名</p>"
          }
        ]
      }
    },
    "success": {
      "fields": {
        "Success 200": [
          {
            "group": "Success 200",
            "optional": false,
            "field": "200",
            "description": "<p>OK</p>"
          }
        ]
      }
    },
    "examples": [
      {
        "title": "BIN",
        "content": "返回文件数据",
        "type": "json"
      }
    ],
    "version": "0.0.0",
    "filename": "./resource.go",
    "groupTitle": "resource",
    "name": "Get"
  },
  {
    "type": "",
    "url": "POST /api/v1/resource/resource 新增一个资源",
    "title": "只有管理员能创建类型type为站内新闻的资源",
    "group": "resource",
    "query": [
      {
        "group": "Query",
        "optional": false,
        "field": "userid",
        "description": "<p>int 必填，执行请求的用户id</p>"
      },
      {
        "group": "Query",
        "optional": false,
        "field": "token",
        "description": "<p>string 必填，执行请求的用户token</p>"
      }
    ],
    "parameter": {
      "fields": {
        "Parameter": [
          {
            "group": "Parameter",
            "optional": false,
            "field": "type",
            "description": "<p>string 必填，资源类型，比如地图、线索、模组等，如果是站内新闻，则房间ID不填，type填&quot;站内新闻&quot;</p>"
          },
          {
            "group": "Parameter",
            "optional": false,
            "field": "rid",
            "description": "<p>string 选填，房间id</p>"
          },
          {
            "group": "Parameter",
            "optional": false,
            "field": "filename",
            "description": "<p>string 必填，资源文件名</p>"
          },
          {
            "group": "Parameter",
            "optional": false,
            "field": "md5",
            "description": "<p>string 必填，资源md5</p>"
          },
          {
            "group": "Parameter",
            "optional": false,
            "field": "resume",
            "description": "<p>string 选填，资源描述</p>"
          },
          {
            "group": "Parameter",
            "optional": false,
            "field": "file",
            "description": "<p>file 选填，文件数据流</p>"
          },
          {
            "group": "Parameter",
            "optional": false,
            "field": "base64File",
            "description": "<p>string 选填，base64文件数据流</p>"
          },
          {
            "group": "Parameter",
            "optional": false,
            "field": "layer_id",
            "description": "<p>int 选填,图层ID</p>"
          },
          {
            "group": "Parameter",
            "optional": false,
            "field": "state",
            "description": "<p>int 选填,槽位ID</p>"
          },
          {
            "group": "Parameter",
            "optional": false,
            "field": "errcode",
            "description": "<p>int 错误代码</p>"
          },
          {
            "group": "Parameter",
            "optional": false,
            "field": "errmsg",
            "description": "<p>string 错误信息</p>"
          },
          {
            "group": "Parameter",
            "optional": false,
            "field": "data",
            "description": "<p>json id:资源ID，md5:资源md5，用这个值可以在后台获取文件数据</p>"
          }
        ]
      }
    },
    "success": {
      "fields": {
        "Success 200": [
          {
            "group": "Success 200",
            "optional": false,
            "field": "200",
            "description": "<p>OK(上传过程会发回progress如：{&quot;progress&quot;: &quot;24%&quot;})</p>"
          }
        ]
      }
    },
    "examples": [
      {
        "title": "json",
        "content": "{\n    \"errcode\":0,\n    \"errmsg\":\"ok\",\n    \"data\":{\"id\":1, \"md5\":\"134123412341234123412234123\"}\n}",
        "type": "json"
      },
      {
        "title": "json",
        "content": "服务器内部错误\n{\n    \"errcode\":30001,\n    \"errmsg\":\"system internal error\"\n}",
        "type": "json"
      },
      {
        "title": "json",
        "content": "用户不合法\n{\n    \"errcode\":30305,\n    \"errmsg\":\"auth failed\"\n}",
        "type": "json"
      },
      {
        "title": "json",
        "content": "参数错误\n{\n    \"errcode\":30002,\n    \"errmsg\":\"invalid parameter\"\n}",
        "type": "json"
      },
      {
        "title": "json",
        "content": "操作数据库失败\n{\n    \"errcode\":30009,\n    \"errmsg\":\"database operation error\"\n}",
        "type": "json"
      }
    ],
    "error": {
      "fields": {
        "Error 4xx": [
          {
            "group": "Error 4xx",
            "optional": false,
            "field": "200",
            "description": "<p>OK</p>"
          }
        ]
      }
    },
    "version": "0.0.0",
    "filename": "./resource.go",
    "groupTitle": "resource",
    "name": "PostApiV1ResourceResource"
  },
  {
    "type": "",
    "url": "PUT /api/v1/resource/resource 修改一个资源",
    "title": "只有管理员能修改类型type为站内新闻的资源;只有管理员和房主才能修改地图",
    "group": "resource",
    "query": [
      {
        "group": "Query",
        "optional": false,
        "field": "userid",
        "description": "<p>int 必填，执行请求的用户id</p>"
      },
      {
        "group": "Query",
        "optional": false,
        "field": "token",
        "description": "<p>string 必填，执行请求的用户token</p>"
      },
      {
        "group": "Query",
        "optional": false,
        "field": "id",
        "description": "<p>int 必填，执行需要修改的资源id</p>"
      }
    ],
    "parameter": {
      "fields": {
        "Parameter": [
          {
            "group": "Parameter",
            "optional": false,
            "field": "type",
            "description": "<p>string 必填，资源类型，比如地图、线索、模组等，如果是站内新闻，则房间ID不填，type填&quot;站内新闻&quot;</p>"
          },
          {
            "group": "Parameter",
            "optional": false,
            "field": "filename",
            "description": "<p>string 必填，资源文件名</p>"
          },
          {
            "group": "Parameter",
            "optional": false,
            "field": "md5",
            "description": "<p>string 必填，资源md5</p>"
          },
          {
            "group": "Parameter",
            "optional": false,
            "field": "resume",
            "description": "<p>string 选填，资源描述</p>"
          },
          {
            "group": "Parameter",
            "optional": false,
            "field": "file",
            "description": "<p>file 选填，文件数据流</p>"
          },
          {
            "group": "Parameter",
            "optional": false,
            "field": "base64File",
            "description": "<p>string 选填，base64文件数据流</p>"
          },
          {
            "group": "Parameter",
            "optional": false,
            "field": "layer_id",
            "description": "<p>int 选填，图层id</p>"
          },
          {
            "group": "Parameter",
            "optional": false,
            "field": "state",
            "description": "<p>int 选填，槽id</p>"
          },
          {
            "group": "Parameter",
            "optional": false,
            "field": "errcode",
            "description": "<p>int 错误代码</p>"
          },
          {
            "group": "Parameter",
            "optional": false,
            "field": "errmsg",
            "description": "<p>string 错误信息</p>"
          },
          {
            "group": "Parameter",
            "optional": false,
            "field": "data",
            "description": "<p>json id:资源ID，md5:资源md5，用这个值可以在后台获取文件数据</p>"
          }
        ]
      }
    },
    "success": {
      "fields": {
        "Success 200": [
          {
            "group": "Success 200",
            "optional": false,
            "field": "200",
            "description": "<p>OK(上传过程会发回progress如：{&quot;progress&quot;: &quot;24%&quot;})</p>"
          }
        ]
      }
    },
    "examples": [
      {
        "title": "json",
        "content": "{\n    \"errcode\":0,\n    \"errmsg\":\"ok\",\n    \"data\":{\"id\":1, \"md5\":\"134123412341234123412234123\"}\n}",
        "type": "json"
      },
      {
        "title": "json",
        "content": "服务器内部错误\n{\n    \"errcode\":30001,\n    \"errmsg\":\"system internal error\"\n}",
        "type": "json"
      },
      {
        "title": "json",
        "content": "用户不合法\n{\n    \"errcode\":30305,\n    \"errmsg\":\"auth failed\"\n}",
        "type": "json"
      },
      {
        "title": "json",
        "content": "参数错误\n{\n    \"errcode\":30002,\n    \"errmsg\":\"invalid parameter\"\n}",
        "type": "json"
      },
      {
        "title": "json",
        "content": "操作数据库失败\n{\n    \"errcode\":30009,\n    \"errmsg\":\"database operation error\"\n}",
        "type": "json"
      }
    ],
    "error": {
      "fields": {
        "Error 4xx": [
          {
            "group": "Error 4xx",
            "optional": false,
            "field": "200",
            "description": "<p>OK</p>"
          }
        ]
      }
    },
    "version": "0.0.0",
    "filename": "./resource.go",
    "groupTitle": "resource",
    "name": "PutApiV1ResourceResource"
  },
  {
    "type": "",
    "url": "DELETE",
    "title": "/api/v1/rolecard/info 删除一个技能/武器/职业/角色卡(每个人只能删除自己的,管理员可以删除公用及其他人的)",
    "group": "rolecard",
    "query": [
      {
        "group": "Query",
        "optional": false,
        "field": "userid",
        "description": "<p>int 必填，执行请求的用户id</p>"
      },
      {
        "group": "Query",
        "optional": false,
        "field": "token",
        "description": "<p>string 必填，执行请求的用户token</p>"
      },
      {
        "group": "Query",
        "optional": false,
        "field": "name",
        "description": "<p>string 必填，请求的表名称，比如角色卡：rolecard；技能：skill; 武器:weapon; 职业：job；NPC卡：npccard； 技能子选项: skselect 不同表决定了body数据中的字段</p>"
      }
    ],
    "parameter": {
      "fields": {
        "Parameter": [
          {
            "group": "Parameter",
            "optional": false,
            "field": "ids",
            "description": "<p>array 必填，要删除的id集</p>"
          },
          {
            "group": "Parameter",
            "optional": false,
            "field": "errcode",
            "description": "<p>int 错误代码</p>"
          },
          {
            "group": "Parameter",
            "optional": false,
            "field": "errmsg",
            "description": "<p>string 错误信息</p>"
          }
        ]
      }
    },
    "examples": [
      {
        "title": "json",
        "content": "{\n    \"ids\":[\"123456789012345678901234\",\"123456789012345678901235\"]\n}",
        "type": "json"
      },
      {
        "title": "json",
        "content": "{\n    \"errcode\":0,\n    \"errmsg\":\"ok\"\n}",
        "type": "json"
      },
      {
        "title": "json",
        "content": "参数错误\n{\n    \"errcode\":30002,\n    \"errmsg\":\"invalid parameter\"\n}",
        "type": "json"
      },
      {
        "title": "json",
        "content": "数据库操作失败\n{\n    \"errcode\":30009,\n    \"errmsg\":\"database operation error\"\n}",
        "type": "json"
      }
    ],
    "success": {
      "fields": {
        "Success 200": [
          {
            "group": "Success 200",
            "optional": false,
            "field": "200",
            "description": "<p>OK</p>"
          }
        ]
      }
    },
    "error": {
      "fields": {
        "Error 4xx": [
          {
            "group": "Error 4xx",
            "optional": false,
            "field": "200",
            "description": "<p>OK</p>"
          }
        ]
      }
    },
    "version": "0.0.0",
    "filename": "./rolecard.go",
    "groupTitle": "rolecard",
    "name": "Delete"
  },
  {
    "type": "",
    "url": "GET",
    "title": "/api/v1/rolecard/count 获取记录数量",
    "group": "rolecard",
    "query": [
      {
        "group": "Query",
        "optional": false,
        "field": "userid",
        "description": "<p>int\t必填，用户ID</p>"
      },
      {
        "group": "Query",
        "optional": false,
        "field": "token",
        "description": "<p>string\t必填，登录token</p>"
      },
      {
        "group": "Query",
        "optional": false,
        "field": "name",
        "description": "<p>string 必填，请求的表名称，比如角色卡：rolecard；技能：skill; 武器:weapon; 职业：job； 技能子选项: skselect 不同表决定了body数据中的字段</p>"
      },
      {
        "group": "Query",
        "optional": false,
        "field": "flag",
        "description": "<p>int  选填，0：查询到自己的和公用的； 1：只查询公用的；2：只查询自己的</p>"
      },
      {
        "group": "Query",
        "optional": false,
        "field": "filter",
        "description": "<p>string 选填，过滤的关键词：例如{&quot;status&quot;:&quot;idle&quot;,&quot;machineid&quot;:1}</p>"
      }
    ],
    "success": {
      "fields": {
        "Success 200": [
          {
            "group": "Success 200",
            "optional": false,
            "field": "200",
            "description": "<p>OK</p>"
          }
        ]
      }
    },
    "parameter": {
      "fields": {
        "Parameter": [
          {
            "group": "Parameter",
            "optional": false,
            "field": "errcode",
            "description": "<p>int     错误代码</p>"
          },
          {
            "group": "Parameter",
            "optional": false,
            "field": "errmsg",
            "description": "<p>string  错误信息</p>"
          },
          {
            "group": "Parameter",
            "optional": false,
            "field": "data",
            "description": "<p>int     查询到的记录数量</p>"
          }
        ]
      }
    },
    "examples": [
      {
        "title": "json",
        "content": "{\n\t   \"errcode\": 0,\n\t   \"errmsg\": \"ok\",\n\t   \"data\": 10\n}",
        "type": "json"
      },
      {
        "title": "json",
        "content": "接口所需要的参数有误\n{\"errcode\":30002,\"errmsg\":\"invalid parameter\"}",
        "type": "json"
      },
      {
        "title": "json",
        "content": "操作数据库出现错误\n{\"errcode\":30001,\"errmsg\":\"system internal error\"}",
        "type": "json"
      }
    ],
    "error": {
      "fields": {
        "Error 4xx": [
          {
            "group": "Error 4xx",
            "optional": false,
            "field": "200",
            "description": "<p>OK</p>"
          }
        ]
      }
    },
    "version": "0.0.0",
    "filename": "./rolecard.go",
    "groupTitle": "rolecard",
    "name": "Get"
  },
  {
    "type": "",
    "url": "GET",
    "title": "/api/v1/rolecard/GetNpcById 根据_id获取npccard",
    "group": "rolecard",
    "query": [
      {
        "group": "Query",
        "optional": false,
        "field": "token",
        "description": "<p>string 必填，执行请求的用户token</p>"
      },
      {
        "group": "Query",
        "optional": false,
        "field": "name",
        "description": "<p>string 必填，填npccard</p>"
      },
      {
        "group": "Query",
        "optional": false,
        "field": "ids",
        "description": "<p>string 选填，_id数组</p>"
      }
    ],
    "success": {
      "fields": {
        "Success 200": [
          {
            "group": "Success 200",
            "optional": false,
            "field": "200",
            "description": "<p>OK</p>"
          }
        ]
      }
    },
    "parameter": {
      "fields": {
        "Parameter": [
          {
            "group": "Parameter",
            "optional": false,
            "field": "errcode",
            "description": "<p>int 错误代码</p>"
          },
          {
            "group": "Parameter",
            "optional": false,
            "field": "errmsg",
            "description": "<p>string 错误信息</p>"
          }
        ]
      }
    },
    "examples": [
      {
        "title": "json",
        "content": "{\n    \"data\":\"\"...\n}",
        "type": "json"
      }
    ],
    "version": "0.0.0",
    "filename": "./rolecard.go",
    "groupTitle": "rolecard",
    "name": "Get"
  },
  {
    "type": "",
    "url": "POST",
    "title": "/api/v1/rolecard/get 按特征查询某个一个技能/武器/职业/角色卡（每个人最多只能查到自己的和公用的）",
    "group": "rolecard",
    "query": [
      {
        "group": "Query",
        "optional": false,
        "field": "userid",
        "description": "<p>int 必填，执行请求的用户id</p>"
      },
      {
        "group": "Query",
        "optional": false,
        "field": "token",
        "description": "<p>string 必填，执行请求的用户token</p>"
      },
      {
        "group": "Query",
        "optional": false,
        "field": "name",
        "description": "<p>string 必填，请求的表名称，比如角色卡：rolecard；技能：skill; 武器:weapon; 职业：job； 技能子选项: skselect 不同表决定了body数据中的字段</p>"
      },
      {
        "group": "Query",
        "optional": false,
        "field": "flag",
        "description": "<p>int  选填，0：查询到自己的和公用的； 1：只查询公用的；2：只查询自己的</p>"
      },
      {
        "group": "Query",
        "optional": false,
        "field": "rid",
        "description": "<p>string  选填，房间id，如果name是rolecard的话，该字段用于房主可以查看玩家的角色卡</p>"
      }
    ],
    "parameter": {
      "fields": {
        "Parameter": [
          {
            "group": "Parameter",
            "optional": false,
            "field": "...",
            "description": "<p>json 必填，查询的特征，需要是json，且相应的表有对应的字段</p>"
          },
          {
            "group": "Parameter",
            "optional": false,
            "field": "errcode",
            "description": "<p>int 错误代码</p>"
          },
          {
            "group": "Parameter",
            "optional": false,
            "field": "errmsg",
            "description": "<p>string 错误信息</p>"
          },
          {
            "group": "Parameter",
            "optional": false,
            "field": "data",
            "description": "<p>json 查询到的</p>"
          }
        ]
      }
    },
    "examples": [
      {
        "title": "json",
        "content": "下面查询name为\"会计\"的技能\n{\n    \"name\":\"会计\"\n}",
        "type": "json"
      },
      {
        "title": "json",
        "content": "{\n    \"errcode\":0,\n    \"errmsg\":\"ok\"\n}",
        "type": "json"
      },
      {
        "title": "json",
        "content": "参数错误\n{\n    \"errcode\":30002,\n    \"errmsg\":\"invalid parameter\"\n}",
        "type": "json"
      },
      {
        "title": "json",
        "content": "数据库操作失败\n{\n    \"errcode\":30009,\n    \"errmsg\":\"database operation error\"\n}",
        "type": "json"
      }
    ],
    "success": {
      "fields": {
        "Success 200": [
          {
            "group": "Success 200",
            "optional": false,
            "field": "200",
            "description": "<p>OK</p>"
          }
        ]
      }
    },
    "error": {
      "fields": {
        "Error 4xx": [
          {
            "group": "Error 4xx",
            "optional": false,
            "field": "200",
            "description": "<p>OK</p>"
          }
        ]
      }
    },
    "version": "0.0.0",
    "filename": "./rolecard.go",
    "groupTitle": "rolecard",
    "name": "Post"
  },
  {
    "type": "",
    "url": "POST",
    "title": "/api/v1/rolecard/share 按特征查询某个一个技能/武器/职业/角色卡（不检验userid和token）",
    "group": "rolecard",
    "query": [
      {
        "group": "Query",
        "optional": false,
        "field": "name",
        "description": "<p>string 必填，请求的表名称，比如角色卡：rolecard(默认)；技能：skill; 武器:weapon; 职业：job； 技能子选项: skselect 不同表决定了body数据中的字段</p>"
      }
    ],
    "parameter": {
      "fields": {
        "Parameter": [
          {
            "group": "Parameter",
            "optional": false,
            "field": "...",
            "description": "<p>json 必填，查询的特征，需要是json，且相应的表有对应的字段</p>"
          },
          {
            "group": "Parameter",
            "optional": false,
            "field": "errcode",
            "description": "<p>int 错误代码</p>"
          },
          {
            "group": "Parameter",
            "optional": false,
            "field": "errmsg",
            "description": "<p>string 错误信息</p>"
          },
          {
            "group": "Parameter",
            "optional": false,
            "field": "data",
            "description": "<p>json 查询到的信息</p>"
          }
        ]
      }
    },
    "examples": [
      {
        "title": "json",
        "content": "下面查询name为\"会计\"的技能\n{\n    \"name\":\"会计\"\n}",
        "type": "json"
      },
      {
        "title": "json",
        "content": "{\n\t\"data\":{...}\n    \"errcode\":0,\n    \"errmsg\":\"ok\"\n}",
        "type": "json"
      },
      {
        "title": "json",
        "content": "参数错误\n{\n    \"errcode\":30002,\n    \"errmsg\":\"invalid parameter\"\n}",
        "type": "json"
      },
      {
        "title": "json",
        "content": "数据库操作失败\n{\n    \"errcode\":30009,\n    \"errmsg\":\"database operation error\"\n}",
        "type": "json"
      }
    ],
    "success": {
      "fields": {
        "Success 200": [
          {
            "group": "Success 200",
            "optional": false,
            "field": "200",
            "description": "<p>OK</p>"
          }
        ]
      }
    },
    "error": {
      "fields": {
        "Error 4xx": [
          {
            "group": "Error 4xx",
            "optional": false,
            "field": "200",
            "description": "<p>OK</p>"
          }
        ]
      }
    },
    "version": "0.0.0",
    "filename": "./rolecard.go",
    "groupTitle": "rolecard",
    "name": "Post"
  },
  {
    "type": "",
    "url": "POST /api/v1/rolecard/info 新增一个技能/武器/职业/角色卡(body参数的userid如果不传，则填为跟query参数的userid一样)",
    "title": "只有管理员能创建公用的,只有管理员能为他人创建",
    "group": "rolecard",
    "query": [
      {
        "group": "Query",
        "optional": false,
        "field": "userid",
        "description": "<p>int 必填，执行请求的用户id</p>"
      },
      {
        "group": "Query",
        "optional": false,
        "field": "token",
        "description": "<p>string 必填，执行请求的用户token</p>"
      },
      {
        "group": "Query",
        "optional": false,
        "field": "name",
        "description": "<p>string 必填，请求的表名称，比如角色卡：rolecard；NPC卡：npccard；技能：skill; 武器:weapon; 职业：job； 技能子选项: skselect 不同表决定了body数据中的字段</p>"
      },
      {
        "group": "Query",
        "optional": false,
        "field": "rg",
        "description": "<p>string 选填，如果name为rolecard时，则为必填，选择角色卡的规则，比如coc7</p>"
      }
    ],
    "examples": [
      {
        "title": "json",
        "content": "如果为skill：\n\tname string\n\tini string\n\tgrow string\n\tpro string\n\tinterest string\n\ttotal string\n\tnum int\n\tkind string\n\tbz int\n\tintroduce string\n\tsub string  子项名称，如kexue\n\tuserid int64  所属用户，如果是公用的，则填0\n如果为job：\n\tvalue string\n\tname string\n\tskills []json  技能集合,和skill表字段一样\n\text int 自定义技能数量\n\tfourt []json 具有如下字段:\n\t\tonum string\n\t\tokind string\n\t\ttnum string\n\t\ttkind string\n\t\thnum string\n\t\thkind string\n\t\tfnum string\n\t\tfkind string\n\tintro []json  介绍说明，具有下面三个字段：\n\t\thonesty string\n\t\tpropoint string\n\t\tproskill\n\tpro [8]int 影响系数，8个值的列表\n\tuserid int64  所属用户，如果是公用的，则填0\n如果为weapons：\n\tname string\n\tskill string\n\tdam string\n\ttho string\n\trange string\n\tround string\n\tnum string\n\tprice string\n\terr int\n\ttime string\n\ttype string   类型，比如cg：常规，sq：手枪等\n\tuserid int64  所属用户，如果是公用的，则填0\n如果为rolecard\n\tjob json 职业信息，和job表字段一样\n\tbz []json 技能信息，和skill字段一样\n\tjobwt []string\n\tselskval []string\n\tname json  有如下字段\n\t\ttouxiang     (可以直接把角色卡的头像的文件内容保存在这里，比如base64)\n\t\tjob\n\t\tjobval\n\t\tplayer\n\t\tchartname\n\t\ttime\n\t\tages\n\t\tsex\n\t\taddress\n\t\thometown\n\tattribute json 有如下字段\n\t\tstr,\n\t\tcon,\n\t\tsiz,\n\t\tdex,\n\t\tapp,\n\t\tint1,\n\t\tpow,\n\t\tedu,\n\t\tluck,\n\t\tmov,\n\t\tbuild,\n\t\tdb\n\thp json 有如下字段\n\t\thave string\n\t\ttotal string\n\tmp  json 有如下字段\n\t\thave string\n\t\ttotal string\n\tsan json 有如下字段\n\t\thave string\n\t\ttotal string\n\tweapons  []string\n\tthings []string\n\tmoney []string\n\tstory json 具有如下字段\n\t\tmiaoshu string\n\t\txinnian string\n\t\tzyzr string\n\t\tfeifanzd string\n\t\tbgzw string\n\t\ttedian string\n\t\tbahen string\n\t\tkongju string\n\t\tstory string\n\tmore json 有如下字段\n\t\tjingli string\n\t\thuoban string\n\t\tkesulu string\n\tdrawpic\n\t\tattrpic []json\n\t\tnamepic []json\n\t\tskillspic []json\n\t\tthingspic []json\n\t\tstorypic []json\n\t\tstory string\n\t\ttouxiang string\n\tmind string\n\thealth string\n\ttouniang string\n\tchartid     //这个字段由后端自动分配，创建成果后会返回#格式为规则简称_顺延卡号,比如coc7_2427,代表着这是一张coc7版规则下的角色卡，角色卡号是2427。如果玩家把卡转移转移给别人，自己的卡不会消失，同时别人会生成一张coc7_2846的角色卡，简单来说就是卡的赠予之后，角色卡号会更新，单个卡号就对应一个用户就好\"",
        "type": "json"
      },
      {
        "title": "json",
        "content": "{\n    \"errcode\":0,\n\t\"errmsg\":\"ok\",\n\t\"data\":{\n\t\t...\n\t}\n}",
        "type": "json"
      },
      {
        "title": "json",
        "content": "参数错误\n{\n    \"errcode\":30002,\n    \"errmsg\":\"invalid parameter\"\n}",
        "type": "json"
      },
      {
        "title": "json",
        "content": "数据库操作失败\n{\n    \"errcode\":30009,\n    \"errmsg\":\"database operation error\"\n}",
        "type": "json"
      }
    ],
    "success": {
      "fields": {
        "Success 200": [
          {
            "group": "Success 200",
            "optional": false,
            "field": "200",
            "description": "<p>OK</p>"
          }
        ]
      }
    },
    "parameter": {
      "fields": {
        "Parameter": [
          {
            "group": "Parameter",
            "optional": false,
            "field": "errcode",
            "description": "<p>int 错误代码</p>"
          },
          {
            "group": "Parameter",
            "optional": false,
            "field": "errmsg",
            "description": "<p>string 错误信息</p>"
          },
          {
            "group": "Parameter",
            "optional": false,
            "field": "data",
            "description": "<p>json 新建立的信息，如果是rolecard，则有chartid字段表示分配到的id，其他的如skill等用_id作为唯一标识</p>"
          }
        ]
      }
    },
    "error": {
      "fields": {
        "Error 4xx": [
          {
            "group": "Error 4xx",
            "optional": false,
            "field": "200",
            "description": "<p>OK</p>"
          }
        ]
      }
    },
    "version": "0.0.0",
    "filename": "./rolecard.go",
    "groupTitle": "rolecard",
    "name": "PostApiV1RolecardInfoBodyUseridQueryUserid"
  },
  {
    "type": "",
    "url": "PUT",
    "title": "/api/v1/rolecard/npcUpdate 修改npccard",
    "group": "rolecard",
    "query": [
      {
        "group": "Query",
        "optional": false,
        "field": "userid",
        "description": "<p>int 必填，执行请求的用户id</p>"
      },
      {
        "group": "Query",
        "optional": false,
        "field": "token",
        "description": "<p>string 必填，执行请求的用户token</p>"
      },
      {
        "group": "Query",
        "optional": false,
        "field": "name",
        "description": "<p>string 必填，填npccard</p>"
      }
    ],
    "parameter": {
      "fields": {
        "Parameter": [
          {
            "group": "Parameter",
            "optional": false,
            "field": "old",
            "description": "<p>json</p>"
          },
          {
            "group": "Parameter",
            "optional": false,
            "field": "new",
            "description": "<p>json</p>"
          },
          {
            "group": "Parameter",
            "optional": false,
            "field": "errcode",
            "description": "<p>int 错误代码</p>"
          },
          {
            "group": "Parameter",
            "optional": false,
            "field": "errmsg",
            "description": "<p>string 错误信息</p>"
          }
        ]
      }
    },
    "examples": [
      {
        "title": "json",
        "content": "{\n\t\"old\":{\n\t\t\"title\":\"aaa\"\n\t},\n\t\"new\":{\n\t\t\"title\":\"bbb\"\n\t}\n}",
        "type": "json"
      },
      {
        "title": "json",
        "content": "{\n    \"errcode\":0,\n    \"errmsg\":\"ok\"\n}",
        "type": "json"
      }
    ],
    "success": {
      "fields": {
        "Success 200": [
          {
            "group": "Success 200",
            "optional": false,
            "field": "200",
            "description": "<p>OK</p>"
          }
        ]
      }
    },
    "version": "0.0.0",
    "filename": "./rolecard.go",
    "groupTitle": "rolecard",
    "name": "Put"
  },
  {
    "type": "",
    "url": "PUT /api/v1/rolecard/update 按特征修改某个一个技能/武器/职业/角色卡（每个人最多只能修改到自己的）",
    "title": "当有在房间的角色卡修改时，会通过websocket通知房间所有玩家, websocket cmd为updateOneRoleCard，extend中有修改后的整个角色卡的信息",
    "group": "rolecard",
    "query": [
      {
        "group": "Query",
        "optional": false,
        "field": "userid",
        "description": "<p>int\t必填，用户ID</p>"
      },
      {
        "group": "Query",
        "optional": false,
        "field": "token",
        "description": "<p>string\t必填，登录token</p>"
      },
      {
        "group": "Query",
        "optional": false,
        "field": "name",
        "description": "<p>string 必填，请求的表名称，比如角色卡：rolecard；技能：skill; 武器:weapon; 职业：job； 技能子选项: skselect 不同表决定了body数据中的字段</p>"
      },
      {
        "group": "Query",
        "optional": false,
        "field": "position",
        "description": "<p>string  取上一次查询结果数组中第一个元素的“_id”作为参数</p>"
      },
      {
        "group": "Query",
        "optional": false,
        "field": "page",
        "description": "<p>int  进行操作所在的页</p>"
      },
      {
        "group": "Query",
        "optional": false,
        "field": "nextpage",
        "description": "<p>int  需要跳转的页</p>"
      },
      {
        "group": "Query",
        "optional": false,
        "field": "pagesize",
        "description": "<p>int  必填，需要返回的数据条数</p>"
      },
      {
        "group": "Query",
        "optional": false,
        "field": "father",
        "description": "<p>string 选填，可以是product、order等（表示把这个工具放到哪一个对象下面），如果不填则为product</p>"
      },
      {
        "group": "Query",
        "optional": false,
        "field": "filter",
        "description": "<p>string 选填，过滤的关键词：例如name=rolecard，filter为{&quot;status&quot;:&quot;idle&quot;,&quot;userid&quot;:1}可以查到用户1的所拥有的角色卡</p>"
      },
      {
        "group": "Query",
        "optional": false,
        "field": "userid",
        "description": "<p>int 必填，执行请求的用户id</p>"
      },
      {
        "group": "Query",
        "optional": false,
        "field": "token",
        "description": "<p>string 必填，执行请求的用户token</p>"
      },
      {
        "group": "Query",
        "optional": false,
        "field": "name",
        "description": "<p>string 必填，请求的表名称，比如角色卡：rolecard；技能：skill; 武器:weapon; 职业：job； 技能子选项: skselect 不同表决定了body数据中的字段</p>"
      },
      {
        "group": "Query",
        "optional": false,
        "field": "filter",
        "description": "<p>string 选填，匹配的关键词, 即要修改哪个，比如要修改rolecard的chartid=coc7_1234的角色卡，则filter={&quot;chartid&quot;:&quot;coc7_1234&quot;}</p>"
      },
      {
        "group": "Query",
        "optional": false,
        "field": "rid",
        "description": "<p>string 选填，角色卡所在房间。一般情况下，修改用户只能修改自己拥有的角色卡，但当该角色卡在房间中，则该房间的房主也可以修改该角色卡；另外当该字段有填，则修改后会把角色卡信息推送到该房间的websocket上</p>"
      }
    ],
    "success": {
      "fields": {
        "Success 200": [
          {
            "group": "Success 200",
            "optional": false,
            "field": "200",
            "description": "<p>OK</p>"
          }
        ]
      }
    },
    "parameter": {
      "fields": {
        "Parameter": [
          {
            "group": "Parameter",
            "optional": false,
            "field": "errcode",
            "description": "<p>int     错误代码</p>"
          },
          {
            "group": "Parameter",
            "optional": false,
            "field": "errmsg",
            "description": "<p>string  错误信息</p>"
          },
          {
            "group": "Parameter",
            "optional": false,
            "field": "data",
            "description": "<p>int     查询到的工具记录数量</p>"
          }
        ]
      }
    },
    "examples": [
      {
        "title": "json",
        "content": "{\n\t   \"errcode\": 0,\n\t   \"errmsg\": \"ok\",\n\t   \"data\": [\n            {\n\t\t\t\t//...\n            }\n        ]\n}",
        "type": "json"
      },
      {
        "title": "json",
        "content": "接口所需要的参数有误\n{\"errcode\":30002,\"errmsg\":\"invalid parameter\"}",
        "type": "json"
      },
      {
        "title": "json",
        "content": "操作数据库出现错误\n{\"errcode\":30001,\"errmsg\":\"system internal error\"}\n/\n\nfunc GetList(c *gin.Context) {\n\t//if Check(c) == false {\n\t//\tR.RJson(c, \"AUTH_FAILED\")\n\t//\treturn\n\t//}\n\t\n\tname := c.Query(\"name\")\n\tif name == \"\" {\n\t\tR.RJson(c, \"INVALID_PARAM\")\n\t\treturn\n\t} else if name != \"job\" && name != \"skill\" && name != \"rolecard\" && name != \"skselect\" && name != \"weapon\" && name != \"npccard\" && name != \"article\" && name != \"invitation\" && name != \"article_comment\" && name != \"invitation_comment\" && name != \"report\" && name != \"manager_report\" {\n\t\tR.RJson(c, \"INVALID_PARAM\")\n\t\treturn\n\t}\n\t\n\t// index := c.Query(\"index\") // []string\n\tposition := c.Query(\"position\")\n\tpage_str := c.Query(\"page\")\n\tnextpage_str := c.Query(\"nextpage\")\n\tpagesize_str := c.Query(\"pagesize\")\n\tfullText := c.Query(\"full_text\")\n\tsort := c.Query(\"sort\")\n\tif sort == \"\" {\n\t\tsort = \"-create_time\"\n\t}\n\t\n\tif name == \"\" {\n\t\tR.RJson(c, \"INVALID_PARAM\")\n\t\treturn\n\t}\n\t\n\tpage, err := strconv.Atoi(page_str)\n\tif err != nil {\n\t\tR.RJson(c, \"INVALID_PARAM\")\n\t\treturn\n\t}\n\tnextpage, err := strconv.Atoi(nextpage_str)\n\tif err != nil {\n\t\tR.RJson(c, \"INVALID_PARAM\")\n\t\treturn\n\t}\n\tpagesize, err := strconv.Atoi(pagesize_str)\n\tif err != nil {\n\t\tR.RJson(c, \"INVALID_PARAM\")\n\t\treturn\n\t}\n\t\n\tneedSearchCharMap := make(map[string]interface{})\n\tdb := Mgodb.C(name)\n\tif fullText != \"\" {\n\t\t//err = db.EnsureIndexKey(\"$text:$**\")//{\"name\":{$regex:\"83\"}}\n\t\t//if err != nil {\n\t\t//\tfmt.Printf(\"EnsureIndexKey fail err=%v\",err)\n\t\t//}\n\t\tneedSearchCharMap[\"title\"] = bson.M{\n\t\t\t\"$regex\": fullText,\n\t\t}\n\t}\n\tfilter := c.Query(\"filter\")\n\tif filter != \"\" {\n\t\tf := make(map[string]interface{})\n\t\terr := json.Unmarshal([]byte(filter), &f)\n\t\tif err == nil {\n\t\t\tswitch name {\n\t\t\tcase \"invitation_comment\", \"article_comment\":\n\t\t\t\tbussID, ok := f[\"buss_id\"].(string)\n\t\t\t\tif !ok {\n\t\t\t\t\tR.RJson(c, \"INVALID_PARAM\")\n\t\t\t\t\treturn\n\t\t\t\t}\n\t\t\t\tf[\"buss_id\"] = bson.ObjectIdHex(bussID)\n\t\t\t}\n\t\t\tfor k, v := range f {\n\t\t\t\tif k == \"_id\" {\n\t\t\t\t\tid_key_str, ok := v.(string)\n\t\t\t\t\tif !ok {\n\t\t\t\t\t\tR.RJson(c, \"INVALID_PARAM\")\n\t\t\t\t\t\treturn\n\t\t\t\t\t}\n\t\t\t\t\tneedSearchCharMap[k] = bson.ObjectIdHex(id_key_str)\n\t\t\t\t} else {\n\t\t\t\t\tneedSearchCharMap[k] = v\n\t\t\t\t}\n\t\t\t}\n\t\t}\n\t}\n\tfmt.Printf(\"needSearchCharMap=%#v\\n\", needSearchCharMap)\n\tvar sampleInfoArr []map[string]interface{}\n\tsortField := make([]string, 0)\n\tvar skip int\n\tif page-nextpage <= 0 {\n\t\tif nextpage-1 < 0 {\n\t\t\tskip = 0\n\t\t} else {\n\t\t\tskip = (nextpage - 1) * pagesize\n\t\t}\n\t\tsortField = append(sortField, sort)\n\t\tif position != \"\" && position != \"0\" {\n\t\t\tneedSearchCharMap[\"_id\"] = map[string]interface{}{\"$lte\": bson.ObjectIdHex(position)}\n\t\t}\n\t\t\n\t\terr = db.Find(needSearchCharMap).\n\t\t\tSkip(skip).\n\t\t\tLimit(pagesize).Sort(sortField...).All(&sampleInfoArr)\n\t} else if page-nextpage > 0 {\n\t\tif nextpage-1 < 0 {\n\t\t\tskip = 0 * pagesize\n\t\t} else {\n\t\t\tskip = (nextpage - 1) * pagesize\n\t\t}\n\t\tsortField = append(sortField, sort)\n\t\tif position != \"\" {\n\t\t\tneedSearchCharMap[\"_id\"] = map[string]interface{}{\"$gt\": bson.ObjectIdHex(position)}\n\t\t}\n\n\t\terr = db.Find(needSearchCharMap).\n\t\t\tSkip(skip).\n\t\t\tLimit(pagesize).Sort(sortField...).All(&sampleInfoArr)\n\t\t\n\t\t//因排序倒置,倒转数组后再发送,统一操作逻辑\n\t\tinfoStart := 0\n\t\tinfoEnd := len(sampleInfoArr) - 1\n\t\tfor infoStart < infoEnd {\n\t\t\tsampleInfoArr[infoStart], sampleInfoArr[infoEnd] =\n\t\t\t\tsampleInfoArr[infoEnd], sampleInfoArr[infoStart]\n\t\t\tinfoStart++\n\t\t\tinfoEnd--\n\t\t}\n\t}\n\t\n\t//end:\n\tif err != nil {\n\t\tfmt.Printf(\"mongo find err=%v\\n\", err)\n\t\tR.RJson(c, \"INTERNAL_ERROR\")\n\t\treturn\n\t}\n\t\n\tR.RData(c, sampleInfoArr)\n}\n\n/*",
        "type": "json"
      },
      {
        "title": "json",
        "content": "{\n\t\"order\":\"123\",\n\t\"step\":\"middle\"\n}",
        "type": "json"
      },
      {
        "title": "json",
        "content": "{\n    \"errcode\":0,\n    \"errmsg\":\"ok\"\n}",
        "type": "json"
      },
      {
        "title": "json",
        "content": "参数错误\n{\n    \"errcode\":30002,\n    \"errmsg\":\"invalid parameter\"\n}",
        "type": "json"
      },
      {
        "title": "json",
        "content": "数据库操作失败\n{\n    \"errcode\":30009,\n    \"errmsg\":\"database operation error\"\n}",
        "type": "json"
      }
    ],
    "error": {
      "fields": {
        "Error 4xx": [
          {
            "group": "Error 4xx",
            "optional": false,
            "field": "200",
            "description": "<p>OK</p>"
          }
        ]
      }
    },
    "version": "0.0.0",
    "filename": "./rolecard.go",
    "groupTitle": "rolecard",
    "name": "PutApiV1RolecardUpdate"
  },
  {
    "type": "",
    "url": "GET",
    "title": "/api/v1/room/roompass 检测房间密码是否正确",
    "group": "room",
    "parameter": {
      "fields": {
        "Parameter": [
          {
            "group": "Parameter",
            "optional": false,
            "field": "userid",
            "description": "<p>int 必填，执行请求的用户id</p>"
          },
          {
            "group": "Parameter",
            "optional": false,
            "field": "token",
            "description": "<p>string 必填，执行请求的用户token</p>"
          },
          {
            "group": "Parameter",
            "optional": false,
            "field": "id",
            "description": "<p>string 选填，房间号</p>"
          },
          {
            "group": "Parameter",
            "optional": false,
            "field": "name",
            "description": "<p>string 选填，房间名称（id和name必填一个，如果两个都填，以id为准）</p>"
          },
          {
            "group": "Parameter",
            "optional": false,
            "field": "password",
            "description": "<p>string 必填，要检查的对应房间的密码</p>"
          }
        ]
      }
    },
    "success": {
      "fields": {
        "Success 200": [
          {
            "group": "Success 200",
            "optional": false,
            "field": "200",
            "description": "<p>OK</p>"
          }
        ]
      }
    },
    "examples": [
      {
        "title": "json",
        "content": "{\n    \"data\": \"true\",\n    \"errcode\": 0,\n    \"errmsg\": \"ok\"\n}",
        "type": "json"
      },
      {
        "title": "json",
        "content": "用户不合法\n{\n    \"errcode\":30305,\n    \"errmsg\":\"auth failed\"\n}",
        "type": "json"
      },
      {
        "title": "json",
        "content": "参数错误\n{\n    \"errcode\":30002,\n    \"errmsg\":\"invalid parameter\"\n}",
        "type": "json"
      },
      {
        "title": "json",
        "content": "房间没找到\n{\n    \"errcode\":30015,\n    \"errmsg\":\"not found\"\n}",
        "type": "json"
      }
    ],
    "error": {
      "fields": {
        "Error 4xx": [
          {
            "group": "Error 4xx",
            "optional": false,
            "field": "200",
            "description": "<p>OK 服务器内部错误 { &quot;errcode&quot;:30001, &quot;errmsg&quot;:&quot;system internal error&quot; }</p>"
          }
        ]
      }
    },
    "version": "0.0.0",
    "filename": "./room.go",
    "groupTitle": "room",
    "name": "Get"
  },
  {
    "type": "",
    "url": "POST",
    "title": "/api/v1/user_ban/ban_user 新增封号用户",
    "group": "user_ban",
    "query": [
      {
        "group": "Query",
        "optional": false,
        "field": "userid",
        "description": "<p>int 必填，执行请求的用户id（只有管理员用户才能封禁,管理员id在config.json中admin处设置）</p>"
      },
      {
        "group": "Query",
        "optional": false,
        "field": "token",
        "description": "<p>string 必填，执行请求的用户token</p>"
      },
      {
        "group": "Query",
        "optional": false,
        "field": "times",
        "description": "<p>int 必填，要封禁的时间，以天为单位</p>"
      },
      {
        "group": "Query",
        "optional": false,
        "field": "uid",
        "description": "<p>string 必填，要封禁用户的id</p>"
      }
    ],
    "examples": [
      {
        "title": "json",
        "content": "{\n    \"data\": {\n        \"uid\": 2,\n        \"suspend\": true,\n        \"starttime\": \"2021-04-26 12:31:14\",\n        \"endtime\": \"2021-05-11 12:31:14\"\n    },\n    \"errcode\": 0,\n    \"errmsg\": \"ok\"\n}",
        "type": "json"
      },
      {
        "title": "json",
        "content": "服务器内部错误\n{\n    \"errcode\":30001,\n    \"errmsg\":\"system internal error\"\n}",
        "type": "json"
      },
      {
        "title": "json",
        "content": "用户不合法\n{\n    \"errcode\":30305,\n    \"errmsg\":\"auth failed\"\n}",
        "type": "json"
      },
      {
        "title": "json",
        "content": "参数错误\n{\n    \"errcode\":30002,\n    \"errmsg\":\"invalid parameter\"\n}",
        "type": "json"
      },
      {
        "title": "json",
        "content": "操作数据库失败\n{\n    \"errcode\":30009,\n    \"errmsg\":\"database operation error\"\n}",
        "type": "json"
      }
    ],
    "error": {
      "fields": {
        "Error 4xx": [
          {
            "group": "Error 4xx",
            "optional": false,
            "field": "200",
            "description": "<p>OK</p>"
          }
        ]
      }
    },
    "parameter": {
      "fields": {
        "Parameter": [
          {
            "group": "Parameter",
            "optional": false,
            "field": "errcode",
            "description": "<p>int 错误代码</p>"
          },
          {
            "group": "Parameter",
            "optional": false,
            "field": "errmsg",
            "description": "<p>string 错误信息</p>"
          }
        ]
      }
    },
    "version": "0.0.0",
    "filename": "./user_ban.go",
    "groupTitle": "user_ban",
    "name": "Post"
  },
  {
    "type": "",
    "url": "GET",
    "title": "/api/v1/collection/detail 获取mongo collection详细",
    "group": "文章帖子",
    "query": [
      {
        "group": "Query",
        "optional": false,
        "field": "userid",
        "description": "<p>int\t必填，用户ID</p>"
      },
      {
        "group": "Query",
        "optional": false,
        "field": "token",
        "description": "<p>string\t必填，登录token</p>"
      },
      {
        "group": "Query",
        "optional": false,
        "field": "name",
        "description": "<p>string 必填，请求的表名称，比如角色卡：rolecard；技能：skill; 武器:weapon; 职业：job； 技能子选项: skselect 不同表决定了body数据中的字段</p>"
      },
      {
        "group": "Query",
        "optional": false,
        "field": "id",
        "description": "<p>string  取上一次查询结果数组中第一个元素的“_id”作为参数</p>"
      },
      {
        "group": "Query",
        "optional": false,
        "field": "filter",
        "description": "<p>string 选填，过滤的关键词：例如name=rolecard，filter为{&quot;status&quot;:&quot;idle&quot;,&quot;userid&quot;:1}可以查到用户1的所拥有的角色卡</p>"
      }
    ],
    "success": {
      "fields": {
        "Success 200": [
          {
            "group": "Success 200",
            "optional": false,
            "field": "200",
            "description": "<p>OK</p>"
          }
        ]
      }
    },
    "parameter": {
      "fields": {
        "Parameter": [
          {
            "group": "Parameter",
            "optional": false,
            "field": "errcode",
            "description": "<p>int     错误代码</p>"
          },
          {
            "group": "Parameter",
            "optional": false,
            "field": "errmsg",
            "description": "<p>string  错误信息</p>"
          },
          {
            "group": "Parameter",
            "optional": false,
            "field": "data",
            "description": "<p>int     查询到的工具记录数量</p>"
          }
        ]
      }
    },
    "examples": [
      {
        "title": "json",
        "content": "{\n\t   \"errcode\": 0,\n\t   \"errmsg\": \"ok\",\n\t   \"data\": [\n            {\n\t\t\t\t//...\n            }\n        ]\n}",
        "type": "json"
      },
      {
        "title": "json",
        "content": "接口所需要的参数有误\n{\"errcode\":30002,\"errmsg\":\"invalid parameter\"}",
        "type": "json"
      },
      {
        "title": "json",
        "content": "操作数据库出现错误\n{\"errcode\":30001,\"errmsg\":\"system internal error\"}",
        "type": "json"
      }
    ],
    "error": {
      "fields": {
        "Error 4xx": [
          {
            "group": "Error 4xx",
            "optional": false,
            "field": "200",
            "description": "<p>OK</p>"
          }
        ]
      }
    },
    "version": "0.0.0",
    "filename": "./article.go",
    "groupTitle": "文章帖子",
    "name": "Get"
  },
  {
    "type": "",
    "url": "GET",
    "title": "/api/v1/report/update 举报审核",
    "group": "文章帖子",
    "query": [
      {
        "group": "Query",
        "optional": false,
        "field": "userid",
        "description": "<p>int\t必填，用户ID</p>"
      },
      {
        "group": "Query",
        "optional": false,
        "field": "token",
        "description": "<p>string\t必填，登录token</p>"
      },
      {
        "group": "Query",
        "optional": false,
        "field": "name",
        "description": "<p>string 必填，请求的表名称 例“report”</p>"
      },
      {
        "group": "Query",
        "optional": false,
        "field": "id",
        "description": "<p>string  举报业务id</p>"
      },
      {
        "group": "Query",
        "optional": false,
        "field": "stat",
        "description": "<p>string 必填 2通过 3不通过 4删除</p>"
      }
    ],
    "success": {
      "fields": {
        "Success 200": [
          {
            "group": "Success 200",
            "optional": false,
            "field": "200",
            "description": "<p>OK</p>"
          }
        ]
      }
    },
    "parameter": {
      "fields": {
        "Parameter": [
          {
            "group": "Parameter",
            "optional": false,
            "field": "errcode",
            "description": "<p>int     错误代码</p>"
          },
          {
            "group": "Parameter",
            "optional": false,
            "field": "errmsg",
            "description": "<p>string  错误信息</p>"
          },
          {
            "group": "Parameter",
            "optional": false,
            "field": "data",
            "description": "<p>int     查询到的工具记录数量</p>"
          }
        ]
      }
    },
    "examples": [
      {
        "title": "json",
        "content": "{\n\t   \"errcode\": 0,\n\t   \"errmsg\": \"ok\",\n\t   \"data\": [\n            {\n\t\t\t\t//...\n            }\n        ]\n}",
        "type": "json"
      },
      {
        "title": "json",
        "content": "接口所需要的参数有误\n{\"errcode\":30002,\"errmsg\":\"invalid parameter\"}",
        "type": "json"
      },
      {
        "title": "json",
        "content": "操作数据库出现错误\n{\"errcode\":30001,\"errmsg\":\"system internal error\"}",
        "type": "json"
      }
    ],
    "error": {
      "fields": {
        "Error 4xx": [
          {
            "group": "Error 4xx",
            "optional": false,
            "field": "200",
            "description": "<p>OK</p>"
          }
        ]
      }
    },
    "version": "0.0.0",
    "filename": "./article.go",
    "groupTitle": "文章帖子",
    "name": "Get"
  },
  {
    "type": "",
    "url": "GET",
    "title": "/api/v1/client_user/set_manager 将前台用户设置成管理员",
    "group": "文章帖子",
    "parameter": {
      "fields": {
        "Parameter": [
          {
            "group": "Parameter",
            "optional": false,
            "field": "userid",
            "description": "<p>int 必填，执行请求的用户id</p>"
          },
          {
            "group": "Parameter",
            "optional": false,
            "field": "token",
            "description": "<p>string 必填，执行请求的用户token</p>"
          },
          {
            "group": "Parameter",
            "optional": false,
            "field": "manager_uid",
            "description": "<p>string 前台管理员用户id</p>"
          }
        ]
      }
    },
    "success": {
      "fields": {
        "Success 200": [
          {
            "group": "Success 200",
            "optional": false,
            "field": "200",
            "description": "<p>OK</p>"
          }
        ]
      }
    },
    "examples": [
      {
        "title": "json",
        "content": "{\n    \"errcode\": 0,\n    \"errmsg\": \"ok\"\n}",
        "type": "json"
      },
      {
        "title": "json",
        "content": "用户不合法\n{\n    \"errcode\":30305,\n    \"errmsg\":\"auth failed\"\n}",
        "type": "json"
      },
      {
        "title": "json",
        "content": "参数错误\n{\n    \"errcode\":30002,\n    \"errmsg\":\"invalid parameter\"\n}",
        "type": "json"
      },
      {
        "title": "json",
        "content": "房间没找到\n{\n    \"errcode\":30015,\n    \"errmsg\":\"not found\"\n}",
        "type": "json"
      }
    ],
    "error": {
      "fields": {
        "Error 4xx": [
          {
            "group": "Error 4xx",
            "optional": false,
            "field": "200",
            "description": "<p>OK 服务器内部错误 { &quot;errcode&quot;:30001, &quot;errmsg&quot;:&quot;system internal error&quot; }</p>"
          }
        ]
      }
    },
    "version": "0.0.0",
    "filename": "./user.go",
    "groupTitle": "文章帖子",
    "name": "Get"
  },
  {
    "type": "",
    "url": "GET",
    "title": "/api/v1/client_user/remove_manager 移除后台管理员用户列表",
    "group": "文章帖子",
    "parameter": {
      "fields": {
        "Parameter": [
          {
            "group": "Parameter",
            "optional": false,
            "field": "userid",
            "description": "<p>int 必填，执行请求的用户id</p>"
          },
          {
            "group": "Parameter",
            "optional": false,
            "field": "token",
            "description": "<p>string 必填，执行请求的用户token</p>"
          },
          {
            "group": "Parameter",
            "optional": false,
            "field": "manager_uid",
            "description": "<p>string 前台管理员用户id</p>"
          }
        ]
      }
    },
    "success": {
      "fields": {
        "Success 200": [
          {
            "group": "Success 200",
            "optional": false,
            "field": "200",
            "description": "<p>OK</p>"
          }
        ]
      }
    },
    "examples": [
      {
        "title": "json",
        "content": "{\n    \"errcode\": 0,\n    \"errmsg\": \"ok\"\n}",
        "type": "json"
      },
      {
        "title": "json",
        "content": "用户不合法\n{\n    \"errcode\":30305,\n    \"errmsg\":\"auth failed\"\n}",
        "type": "json"
      },
      {
        "title": "json",
        "content": "参数错误\n{\n    \"errcode\":30002,\n    \"errmsg\":\"invalid parameter\"\n}",
        "type": "json"
      },
      {
        "title": "json",
        "content": "房间没找到\n{\n    \"errcode\":30015,\n    \"errmsg\":\"not found\"\n}",
        "type": "json"
      }
    ],
    "error": {
      "fields": {
        "Error 4xx": [
          {
            "group": "Error 4xx",
            "optional": false,
            "field": "200",
            "description": "<p>OK 服务器内部错误 { &quot;errcode&quot;:30001, &quot;errmsg&quot;:&quot;system internal error&quot; }</p>"
          }
        ]
      }
    },
    "version": "0.0.0",
    "filename": "./user.go",
    "groupTitle": "文章帖子",
    "name": "Get"
  },
  {
    "type": "",
    "url": "GET",
    "title": "/api/v1/client_user/manager_list 后台查看前台管理员用户列表",
    "group": "文章帖子",
    "parameter": {
      "fields": {
        "Parameter": [
          {
            "group": "Parameter",
            "optional": false,
            "field": "userid",
            "description": "<p>int 必填，执行请求的用户id</p>"
          },
          {
            "group": "Parameter",
            "optional": false,
            "field": "token",
            "description": "<p>string 必填，执行请求的用户token</p>"
          },
          {
            "group": "Parameter",
            "optional": false,
            "field": "page_index",
            "description": "<p>string 选填，页码</p>"
          },
          {
            "group": "Parameter",
            "optional": false,
            "field": "page_size",
            "description": "<p>string 选填，每页数量</p>"
          }
        ]
      }
    },
    "success": {
      "fields": {
        "Success 200": [
          {
            "group": "Success 200",
            "optional": false,
            "field": "200",
            "description": "<p>OK</p>"
          }
        ]
      }
    },
    "examples": [
      {
        "title": "json",
        "content": "{\n    \"data\": {\n        \"list\":[\n         {\"name\":\"xx\",\"id\":1,\"touxiang\":\"xx\"},\n         {\"name\":\"xx\",\"id\":1,\"touxiang\":\"xx\"},\n\n         ],\n         \"total\":2\n    },\n    \"errcode\": 0,\n    \"errmsg\": \"ok\"\n}",
        "type": "json"
      },
      {
        "title": "json",
        "content": "用户不合法\n{\n    \"errcode\":30305,\n    \"errmsg\":\"auth failed\"\n}",
        "type": "json"
      },
      {
        "title": "json",
        "content": "参数错误\n{\n    \"errcode\":30002,\n    \"errmsg\":\"invalid parameter\"\n}",
        "type": "json"
      },
      {
        "title": "json",
        "content": "房间没找到\n{\n    \"errcode\":30015,\n    \"errmsg\":\"not found\"\n}",
        "type": "json"
      }
    ],
    "error": {
      "fields": {
        "Error 4xx": [
          {
            "group": "Error 4xx",
            "optional": false,
            "field": "200",
            "description": "<p>OK 服务器内部错误 { &quot;errcode&quot;:30001, &quot;errmsg&quot;:&quot;system internal error&quot; }</p>"
          }
        ]
      }
    },
    "version": "0.0.0",
    "filename": "./user.go",
    "groupTitle": "文章帖子",
    "name": "Get"
  },
  {
    "type": "",
    "url": "GET",
    "title": "GET /api/v1/client_user/user 获取前台管理员用户信息",
    "group": "文章帖子",
    "query": [
      {
        "group": "Query",
        "optional": false,
        "field": "userid",
        "description": "<p>int 必填，用户id</p>"
      },
      {
        "group": "Query",
        "optional": false,
        "field": "token",
        "description": "<p>string 必填，用户token</p>"
      },
      {
        "group": "Query",
        "optional": false,
        "field": "manager_uid",
        "description": "<p>string 选填，要查询的用户id，id和email必填一项</p>"
      }
    ],
    "success": {
      "fields": {
        "Success 200": [
          {
            "group": "Success 200",
            "optional": false,
            "field": "200",
            "description": "<p>OK</p>"
          }
        ]
      }
    },
    "parameter": {
      "fields": {
        "Parameter": [
          {
            "group": "Parameter",
            "optional": false,
            "field": "errcode",
            "description": "<p>int 错误代码</p>"
          },
          {
            "group": "Parameter",
            "optional": false,
            "field": "errmsg",
            "description": "<p>string 错误信息</p>"
          },
          {
            "group": "Parameter",
            "optional": false,
            "field": "data",
            "description": "<p>json 用户信息</p>"
          }
        ]
      }
    },
    "examples": [
      {
        "title": "json",
        "content": "{\n    \"errcode\":0,\n    \"errmsg\":\"ok\",\n    \"data\":{\n        \"userid\":1,\n        \"name\":\"bob\",\n        \"cnname\":\"bo\",\n        \"email\":\"test@test.com\",\n        \"phone\":88888888,\n        \"im\":12345,\n        \"qq\":54321,\n        \"passwd\":\"*****\",      // 已经设置密码则为*****，没设置密码则为空\n        \"rolecard\":[\"coc7_2002\", \"coc7_4123\"],   //（弃用，用户所拥有的角色卡在rolecard list接口中查看）\n        \"sex\":\"male\",             //性别，male：男、female：女、其他为保密\n        \"touxiang\":\"12345678901234567890123456789012\",    // 头像MD5值\n        \"birthday\":15002939021,   //生日时间戳\n        \"sign\":\"asdfasdfasdfadsf\"    //签名\n    }\n}",
        "type": "json"
      },
      {
        "title": "json",
        "content": "校验token和session失败\n{\n   \"errcode\":30305,\n   \"errmsg\":\"auth failed\"\n}",
        "type": "json"
      },
      {
        "title": "json",
        "content": "无效参数\n{\n    \"errcode\":30002,\n    \"errmsg\":\"invalid parameter\"\n}",
        "type": "json"
      },
      {
        "title": "json",
        "content": "未知的用户\n{\n    \"errcode\":30101,\n    \"errmsg\":\"unknown user\"\n}",
        "type": "json"
      },
      {
        "title": "json",
        "content": "没有权限\n{\n    \"errcode\":30005,\n    \"errmsg\":\"no permission\"\n}",
        "type": "json"
      }
    ],
    "error": {
      "fields": {
        "Error 4xx": [
          {
            "group": "Error 4xx",
            "optional": false,
            "field": "200",
            "description": "<p>OK</p>"
          }
        ]
      }
    },
    "version": "0.0.0",
    "filename": "./user.go",
    "groupTitle": "文章帖子",
    "name": "Get"
  },
  {
    "type": "",
    "url": "GET",
    "title": "/api/v1/manager_report/update  管理员举报审核",
    "group": "文章帖子",
    "query": [
      {
        "group": "Query",
        "optional": false,
        "field": "userid",
        "description": "<p>int\t必填，用户ID</p>"
      },
      {
        "group": "Query",
        "optional": false,
        "field": "token",
        "description": "<p>string\t必填，登录token</p>"
      },
      {
        "group": "Query",
        "optional": false,
        "field": "name",
        "description": "<p>string 必填，请求的表名称 例“report”</p>"
      },
      {
        "group": "Query",
        "optional": false,
        "field": "id",
        "description": "<p>string  举报业务id</p>"
      },
      {
        "group": "Query",
        "optional": false,
        "field": "stat",
        "description": "<p>string 必填 2通过 3不通过 4删除</p>"
      }
    ],
    "success": {
      "fields": {
        "Success 200": [
          {
            "group": "Success 200",
            "optional": false,
            "field": "200",
            "description": "<p>OK</p>"
          }
        ]
      }
    },
    "parameter": {
      "fields": {
        "Parameter": [
          {
            "group": "Parameter",
            "optional": false,
            "field": "errcode",
            "description": "<p>int     错误代码</p>"
          },
          {
            "group": "Parameter",
            "optional": false,
            "field": "errmsg",
            "description": "<p>string  错误信息</p>"
          },
          {
            "group": "Parameter",
            "optional": false,
            "field": "data",
            "description": "<p>int     查询到的工具记录数量</p>"
          }
        ]
      }
    },
    "examples": [
      {
        "title": "json",
        "content": "{\n\t   \"errcode\": 0,\n\t   \"errmsg\": \"ok\",\n\t   \"data\": [\n            {\n\t\t\t\t//...\n            }\n        ]\n}",
        "type": "json"
      },
      {
        "title": "json",
        "content": "接口所需要的参数有误\n{\"errcode\":30002,\"errmsg\":\"invalid parameter\"}",
        "type": "json"
      },
      {
        "title": "json",
        "content": "操作数据库出现错误\n{\"errcode\":30001,\"errmsg\":\"system internal error\"}",
        "type": "json"
      }
    ],
    "error": {
      "fields": {
        "Error 4xx": [
          {
            "group": "Error 4xx",
            "optional": false,
            "field": "200",
            "description": "<p>OK</p>"
          }
        ]
      }
    },
    "version": "0.0.0",
    "filename": "./article.go",
    "groupTitle": "文章帖子",
    "name": "Get"
  }
] });
