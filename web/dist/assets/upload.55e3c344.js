/*! 
 Build based on gin-vue-admin 
 Time : 1679909256000 */
import{O as e,aW as t,q as a,s,bk as l,aR as o,A as n,ai as r,C as i,bS as u,r as c,o as d,a0 as p,w as f,c as v,U as m,aC as y,n as h,J as g,aE as b,I as k,f as w,d as F,ah as R,e as E,av as S,c9 as C,a3 as T,a2 as $,c5 as L,c6 as _,b2 as x,ca as U,cb as P,a4 as O,_ as j,v as D,M as B,ak as q,cc as A,cd as H,D as K,K as M,W,z as X,Q as z,F as N,aI as I,aG as J,aX as G,a5 as Q}from"./index.61dd0a2d.js";import{E as V}from"./progress.d509d1cc.js";const Y=Symbol("uploadContextKey");class Z extends Error{constructor(e,t,a,s){super(e),this.name="UploadAjaxError",this.status=t,this.method=a,this.url=s}}function ee(e,t,a){let s;return s=a.response?`${a.response.error||a.response}`:a.responseText?`${a.responseText}`:`fail to ${t.method} ${e} ${a.status}`,new Z(s,a.status,t.method,e)}const te=["text","picture","picture-card"];let ae=1;const se=()=>Date.now()+ae++,le=a({action:{type:String,default:"#"},headers:{type:s(Object)},method:{type:String,default:"post"},data:{type:Object,default:()=>l({})},multiple:{type:Boolean,default:!1},name:{type:String,default:"file"},drag:{type:Boolean,default:!1},withCredentials:Boolean,showFileList:{type:Boolean,default:!0},accept:{type:String,default:""},type:{type:String,default:"select"},fileList:{type:s(Array),default:()=>l([])},autoUpload:{type:Boolean,default:!0},listType:{type:String,values:te,default:"text"},httpRequest:{type:s(Function),default:a=>{"undefined"==typeof XMLHttpRequest&&e("ElUpload","XMLHttpRequest is undefined");const s=new XMLHttpRequest,l=a.action;s.upload&&s.upload.addEventListener("progress",(e=>{const t=e;t.percent=e.total>0?e.loaded/e.total*100:0,a.onProgress(t)}));const o=new FormData;if(a.data)for(const[e,t]of Object.entries(a.data))Array.isArray(t)?o.append(e,...t):o.append(e,t);o.append(a.filename,a.file,a.file.name),s.addEventListener("error",(()=>{a.onError(ee(l,a,s))})),s.addEventListener("load",(()=>{if(s.status<200||s.status>=300)return a.onError(ee(l,a,s));a.onSuccess(function(e){const t=e.responseText||e.response;if(!t)return t;try{return JSON.parse(t)}catch(a){return t}}(s))})),s.open(a.method,l,!0),a.withCredentials&&"withCredentials"in s&&(s.withCredentials=!0);const n=a.headers||{};if(n instanceof Headers)n.forEach(((e,t)=>s.setRequestHeader(t,e)));else for(const[e,r]of Object.entries(n))t(r)||s.setRequestHeader(e,String(r));return s.send(o),s}},disabled:Boolean,limit:Number}),oe=a({...le,beforeUpload:{type:s(Function),default:o},beforeRemove:{type:s(Function)},onRemove:{type:s(Function),default:o},onChange:{type:s(Function),default:o},onPreview:{type:s(Function),default:o},onSuccess:{type:s(Function),default:o},onProgress:{type:s(Function),default:o},onError:{type:s(Function),default:o},onExceed:{type:s(Function),default:o}}),ne=a({files:{type:s(Array),default:()=>l([])},disabled:{type:Boolean,default:!1},handlePreview:{type:s(Function),default:o},listType:{type:String,values:te,default:"text"}}),re=["onKeydown"],ie=["src"],ue=["onClick"],ce=["onClick"],de=["onClick"],pe=n({name:"ElUploadList"});var fe=j(n({...pe,props:ne,emits:{remove:e=>!!e},setup(e,{emit:t}){const{t:a}=r(),s=i("upload"),l=i("icon"),o=i("list"),n=u(),j=c(!1),D=e=>{t("remove",e)};return(e,t)=>(d(),p(O,{tag:"ul",class:h([g(s).b("list"),g(s).bm("list",e.listType),g(s).is("disabled",g(n))]),name:g(o).b()},{default:f((()=>[(d(!0),v(m,null,y(e.files,(o=>(d(),v("li",{key:o.uid||o.name,class:h([g(s).be("list","item"),g(s).is(o.status),{focusing:j.value}]),tabindex:"0",onKeydown:b((e=>!g(n)&&D(o)),["delete"]),onFocus:t[0]||(t[0]=e=>j.value=!0),onBlur:t[1]||(t[1]=e=>j.value=!1),onClick:t[2]||(t[2]=e=>j.value=!1)},[k(e.$slots,"default",{file:o},(()=>["picture"===e.listType||"uploading"!==o.status&&"picture-card"===e.listType?(d(),v("img",{key:0,class:h(g(s).be("list","item-thumbnail")),src:o.url,alt:""},null,10,ie)):w("v-if",!0),"uploading"===o.status||"picture-card"!==e.listType?(d(),v("div",{key:1,class:h(g(s).be("list","item-info"))},[F("a",{class:h(g(s).be("list","item-name")),onClick:R((t=>e.handlePreview(o)),["prevent"])},[E(g(S),{class:h(g(l).m("document"))},{default:f((()=>[E(g(C))])),_:1},8,["class"]),F("span",{class:h(g(s).be("list","item-file-name"))},T(o.name),3)],10,ue),"uploading"===o.status?(d(),p(g(V),{key:0,type:"picture-card"===e.listType?"circle":"line","stroke-width":"picture-card"===e.listType?6:2,percentage:Number(o.percentage),style:$("picture-card"===e.listType?"":"margin-top: 0.5rem")},null,8,["type","stroke-width","percentage","style"])):w("v-if",!0)],2)):w("v-if",!0),F("label",{class:h(g(s).be("list","item-status-label"))},["text"===e.listType?(d(),p(g(S),{key:0,class:h([g(l).m("upload-success"),g(l).m("circle-check")])},{default:f((()=>[E(g(L))])),_:1},8,["class"])):["picture-card","picture"].includes(e.listType)?(d(),p(g(S),{key:1,class:h([g(l).m("upload-success"),g(l).m("check")])},{default:f((()=>[E(g(_))])),_:1},8,["class"])):w("v-if",!0)],2),g(n)?w("v-if",!0):(d(),p(g(S),{key:2,class:h(g(l).m("close")),onClick:e=>D(o)},{default:f((()=>[E(g(x))])),_:2},1032,["class","onClick"])),w(" Due to close btn only appears when li gets focused disappears after li gets blurred, thus keyboard navigation can never reach close btn"),w(" This is a bug which needs to be fixed "),w(" TODO: Fix the incorrect navigation interaction "),g(n)?w("v-if",!0):(d(),v("i",{key:3,class:h(g(l).m("close-tip"))},T(g(a)("el.upload.deleteTip")),3)),"picture-card"===e.listType?(d(),v("span",{key:4,class:h(g(s).be("list","item-actions"))},[F("span",{class:h(g(s).be("list","item-preview")),onClick:t=>e.handlePreview(o)},[E(g(S),{class:h(g(l).m("zoom-in"))},{default:f((()=>[E(g(U))])),_:1},8,["class"])],10,ce),g(n)?w("v-if",!0):(d(),v("span",{key:0,class:h(g(s).be("list","item-delete")),onClick:e=>D(o)},[E(g(S),{class:h(g(l).m("delete"))},{default:f((()=>[E(g(P))])),_:1},8,["class"])],10,de))],2)):w("v-if",!0)]))],42,re)))),128)),k(e.$slots,"append")])),_:3},8,["class","name"]))}}),[["__file","/home/runner/work/element-plus/element-plus/packages/components/upload/src/upload-list.vue"]]);const ve=a({disabled:{type:Boolean,default:!1}}),me={file:e=>D(e)},ye=["onDrop","onDragover"],he="ElUploadDrag",ge=n({name:he});var be=j(n({...ge,props:ve,emits:me,setup(t,{emit:a}){const s=B(Y);s||e(he,"usage: <el-upload><el-upload-dragger /></el-upload>");const l=i("upload"),o=c(!1),n=u(),r=e=>{if(n.value)return;o.value=!1;const t=Array.from(e.dataTransfer.files),l=s.accept.value;if(!l)return void a("file",t);const r=t.filter((e=>{const{type:t,name:a}=e,s=a.includes(".")?`.${a.split(".").pop()}`:"",o=t.replace(/\/.*$/,"");return l.split(",").map((e=>e.trim())).filter((e=>e)).some((e=>e.startsWith(".")?s===e:/\/\*$/.test(e)?o===e.replace(/\/\*$/,""):!!/^[^/]+\/[^/]+$/.test(e)&&t===e))}));a("file",r)},p=()=>{n.value||(o.value=!0)};return(e,t)=>(d(),v("div",{class:h([g(l).b("dragger"),g(l).is("dragover",o.value)]),onDrop:R(r,["prevent"]),onDragover:R(p,["prevent"]),onDragleave:t[0]||(t[0]=R((e=>o.value=!1),["prevent"]))},[k(e.$slots,"default")],42,ye))}}),[["__file","/home/runner/work/element-plus/element-plus/packages/components/upload/src/upload-dragger.vue"]]);const ke=a({...le,beforeUpload:{type:s(Function),default:o},onRemove:{type:s(Function),default:o},onStart:{type:s(Function),default:o},onSuccess:{type:s(Function),default:o},onProgress:{type:s(Function),default:o},onError:{type:s(Function),default:o},onExceed:{type:s(Function),default:o}}),we=["onKeydown"],Fe=["name","multiple","accept"],Re=n({name:"ElUploadContent",inheritAttrs:!1});var Ee=j(n({...Re,props:ke,setup(e,{expose:t}){const a=e,s=i("upload"),l=u(),o=q({}),n=q(),r=e=>{if(0===e.length)return;const{autoUpload:t,limit:s,fileList:l,multiple:o,onStart:n,onExceed:r}=a;if(s&&l.length+e.length>s)r(e,l);else{o||(e=e.slice(0,1));for(const a of e){const e=a;e.uid=se(),n(e),t&&c(e)}}},c=async e=>{if(n.value.value="",!a.beforeUpload)return m(e);let t;try{t=await a.beforeUpload(e)}catch(l){t=!1}if(!1===t)return void a.onRemove(e);let s=e;t instanceof Blob&&(s=t instanceof File?t:new File([t],e.name,{type:e.type})),m(Object.assign(s,{uid:e.uid}))},m=e=>{const{headers:t,data:s,method:l,withCredentials:n,name:r,action:i,onProgress:u,onSuccess:c,onError:d,httpRequest:p}=a,{uid:f}=e,v={headers:t||{},withCredentials:n,file:e,data:s,method:l,filename:r,action:i,onProgress:t=>{u(t,e)},onSuccess:t=>{c(t,e),delete o.value[f]},onError:t=>{d(t,e),delete o.value[f]}},m=p(v);o.value[f]=m,m instanceof Promise&&m.then(v.onSuccess,v.onError)},y=e=>{const t=e.target.files;t&&r(Array.from(t))},w=()=>{l.value||(n.value.value="",n.value.click())},E=()=>{w()};return t({abort:e=>{A(o.value).filter(e?([t])=>String(e.uid)===t:()=>!0).forEach((([e,t])=>{t instanceof XMLHttpRequest&&t.abort(),delete o.value[e]}))},upload:c}),(e,t)=>(d(),v("div",{class:h([g(s).b(),g(s).m(e.listType),g(s).is("drag",e.drag)]),tabindex:"0",onClick:w,onKeydown:b(R(E,["self"]),["enter","space"])},[e.drag?(d(),p(be,{key:0,disabled:g(l),onFile:r},{default:f((()=>[k(e.$slots,"default")])),_:3},8,["disabled"])):k(e.$slots,"default",{key:1}),F("input",{ref_key:"inputRef",ref:n,class:h(g(s).e("input")),name:e.name,multiple:e.multiple,accept:e.accept,type:"file",onChange:y,onClick:t[0]||(t[0]=R((()=>{}),["stop"]))},null,42,Fe)],42,we))}}),[["__file","/home/runner/work/element-plus/element-plus/packages/components/upload/src/upload-content.vue"]]);const Se="ElUpload",Ce=(a,s)=>{const l=H(a,"fileList",void 0,{passive:!0}),o=e=>l.value.find((t=>t.uid===e.uid));function n(e){var t;null==(t=s.value)||t.abort(e)}return K((()=>a.listType),(e=>{"picture-card"!==e&&"picture"!==e||(l.value=l.value.map((e=>{const{raw:t,url:s}=e;if(!s&&t)try{e.url=URL.createObjectURL(t)}catch(o){a.onError(o,e,l.value)}return e})))})),K(l,(e=>{for(const t of e)t.uid||(t.uid=se()),t.status||(t.status="success")}),{immediate:!0,deep:!0}),{uploadFiles:l,abort:n,clearFiles:function(e=["ready","uploading","success","fail"]){l.value=l.value.filter((t=>!e.includes(t.status)))},handleError:(e,t)=>{const s=o(t);s&&(console.error(e),s.status="fail",l.value.splice(l.value.indexOf(s),1),a.onError(e,s,l.value),a.onChange(s,l.value))},handleProgress:(e,t)=>{const s=o(t);s&&(a.onProgress(e,s,l.value),s.status="uploading",s.percentage=Math.round(e.percent))},handleStart:e=>{t(e.uid)&&(e.uid=se());const s={name:e.name,percentage:0,status:"ready",size:e.size,raw:e,uid:e.uid};if("picture-card"===a.listType||"picture"===a.listType)try{s.url=URL.createObjectURL(e)}catch(o){M(Se,o.message),a.onError(o,s,l.value)}l.value=[...l.value,s],a.onChange(s,l.value)},handleSuccess:(e,t)=>{const s=o(t);s&&(s.status="success",s.response=e,a.onSuccess(e,s,l.value),a.onChange(s,l.value))},handleRemove:async t=>{const s=t instanceof File?o(t):t;s||e(Se,"file to be removed not found");const r=e=>{n(e);const t=l.value;t.splice(t.indexOf(e),1),a.onRemove(e,t),(e=>{var t;(null==(t=e.url)?void 0:t.startsWith("blob:"))&&URL.revokeObjectURL(e.url)})(e)};if(a.beforeRemove){!1!==await a.beforeRemove(s,l.value)&&r(s)}else r(s)},submit:function(){l.value.filter((({status:e})=>"ready"===e)).forEach((({raw:e})=>{var t;return e&&(null==(t=s.value)?void 0:t.upload(e))}))}}},Te=n({name:"ElUpload"});const $e=Q(j(n({...Te,props:oe,setup(e,{expose:t}){const a=e,s=W(),l=u(),o=q(),{abort:n,submit:r,clearFiles:i,uploadFiles:c,handleStart:m,handleError:y,handleRemove:h,handleSuccess:b,handleProgress:F}=Ce(a,o),R=X((()=>"picture-card"===a.listType)),S=X((()=>({...a,fileList:c.value,onStart:m,onProgress:F,onSuccess:b,onError:y,onRemove:h})));return z((()=>{c.value.forEach((({url:e})=>{(null==e?void 0:e.startsWith("blob:"))&&URL.revokeObjectURL(e)}))})),N(Y,{accept:I(a,"accept")}),t({abort:n,submit:r,clearFiles:i,handleStart:m,handleRemove:h}),(e,t)=>(d(),v("div",null,[g(R)&&e.showFileList?(d(),p(fe,{key:0,disabled:g(l),"list-type":e.listType,files:g(c),"handle-preview":e.onPreview,onRemove:g(h)},J({append:f((()=>[E(Ee,G({ref_key:"uploadRef",ref:o},g(S)),{default:f((()=>[g(s).trigger?k(e.$slots,"trigger",{key:0}):w("v-if",!0),!g(s).trigger&&g(s).default?k(e.$slots,"default",{key:1}):w("v-if",!0)])),_:3},16)])),_:2},[e.$slots.file?{name:"default",fn:f((({file:t})=>[k(e.$slots,"file",{file:t})]))}:void 0]),1032,["disabled","list-type","files","handle-preview","onRemove"])):w("v-if",!0),!g(R)||g(R)&&!e.showFileList?(d(),p(Ee,G({key:1,ref_key:"uploadRef",ref:o},g(S)),{default:f((()=>[g(s).trigger?k(e.$slots,"trigger",{key:0}):w("v-if",!0),!g(s).trigger&&g(s).default?k(e.$slots,"default",{key:1}):w("v-if",!0)])),_:3},16)):w("v-if",!0),e.$slots.trigger?k(e.$slots,"default",{key:2}):w("v-if",!0),k(e.$slots,"tip"),!g(R)&&e.showFileList?(d(),p(fe,{key:3,disabled:g(l),"list-type":e.listType,files:g(c),"handle-preview":e.onPreview,onRemove:g(h)},J({_:2},[e.$slots.file?{name:"default",fn:f((({file:t})=>[k(e.$slots,"file",{file:t})]))}:void 0]),1032,["disabled","list-type","files","handle-preview","onRemove"])):w("v-if",!0)]))}}),[["__file","/home/runner/work/element-plus/element-plus/packages/components/upload/src/upload.vue"]]));export{$e as E};
