/*! 
 Build based on gin-vue-admin 
 Time : 1680141068000 */
System.register(["./github-legacy.98276463.js","./date-legacy.fb7d66fc.js","./_plugin-vue_export-helper-legacy.649209c3.js","./index-legacy.33c92641.js"],(function(t,e){"use strict";var a,o,l,c,i,n,m,f,s,r,d,g,u,b=document.createElement("style");return b.textContent=".commit-table[data-v-144ac47f]{background-color:#fff;height:400px}.commit-table-title[data-v-144ac47f]{font-weight:600;margin-bottom:12px}.commit-table .log-item[data-v-144ac47f]{display:flex;justify-content:space-between;margin-top:14px}.commit-table .log-item .key-box[data-v-144ac47f]{justify-content:center}.commit-table .log-item .key[data-v-144ac47f]{display:inline-flex;justify-content:center;align-items:center;width:20px;height:20px;border-radius:50%;background:#F0F2F5;text-align:center;color:rgba(0,0,0,.65)}.commit-table .log-item .key.top[data-v-144ac47f]{background:#314659;color:#fff}.commit-table .log-item .message[data-v-144ac47f]{color:rgba(0,0,0,.65)}.commit-table .log-item .form[data-v-144ac47f]{color:rgba(0,0,0,.65);margin-left:12px}.commit-table .log-item .flex[data-v-144ac47f]{line-height:20px;overflow:hidden;text-overflow:ellipsis;white-space:nowrap}.commit-table .log-item .flex-1[data-v-144ac47f]{flex:1}.commit-table .log-item .flex-2[data-v-144ac47f]{flex:2}.commit-table .log-item .flex-3[data-v-144ac47f]{flex:3}.commit-table .log-item .flex-4[data-v-144ac47f]{flex:4}.commit-table .log-item .flex-5[data-v-144ac47f]{flex:5}\n",document.head.appendChild(b),{setters:[function(t){a=t.C},function(t){o=t.f},function(t){l=t._},function(t){c=t.r,i=t.o,n=t.c,m=t.d,f=t.U,s=t.aC,r=t.n,d=t.a3,g=t.bh,u=t.bi}],execute:function(){var e={class:"commit-table"},b=function(t){return g("data-v-144ac47f"),t=t(),u(),t}((function(){return m("div",{class:"commit-table-title"}," 更新日志 ",-1)})),x={class:"log"},v={class:"flex-1 flex key-box"},p={class:"flex-5 flex message"},h={class:"flex-3 flex form"},y={name:"DashboardTable"};t("default",l(Object.assign(y,{setup:function(t){var l=c(!0),g=c([]);return a(0).then((function(t){var e=t.data;l.value=!1,e.forEach((function(t,e){t.commit.message&&e<10&&g.value.push({from:o(t.commit.author.date,"yyyy-MM-dd"),title:t.commit.author.name,showDayAndMonth:!0,message:t.commit.message})}))})),function(t,a){return i(),n("div",e,[b,m("div",x,[(i(!0),n(f,null,s(g.value,(function(t,e){return i(),n("div",{key:e,class:"log-item"},[m("div",v,[m("span",{class:r(["key",e<3&&"top"])},d(e+1),3)]),m("div",p,d(t.message),1),m("div",h,d(t.from),1)])})),128))])])}}}),[["__scopeId","data-v-144ac47f"]]))}}}));