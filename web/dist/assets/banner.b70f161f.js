/*! 
 Build based on gin-vue-admin 
 Time : 1679909256000 */
import{r as e,a,o as l,c as t,d as o,e as s,w as r,h as i,aE as u,a3 as d,J as n,U as p,aC as c,a0 as m,bm as v,j as f,E as g,b}from"./index.61dd0a2d.js";import{E as y}from"./dialog.a2b2142c.js";import"./overlay.533fa516.js";/* empty css              */import"./tag.3ecd1fed.js";import{E as h,a as j}from"./select.31449efc.js";import"./scrollbar.ef9dc6d6.js";/* empty css               */import{E as _}from"./pagination.7bbc3896.js";import{E as w,a as C}from"./table-column.720e0d1a.js";import"./checkbox.384d52e9.js";/* empty css                */import{E as k,_ as x}from"./index.5c65e0d6.js";import{E as V}from"./popover.4fe0f5fa.js";import{E,a as z}from"./form-item.a3ced827.js";/* empty css               */import{E as U}from"./date-picker.c50427f0.js";import{g as I,d as D,f as A,a as S,c as B,u as q}from"./banner.8a161d42.js";import{f as T,a as K,g as O}from"./format.a3be29dc.js";import"./refs.1f88eb22.js";import"./index.cbcb4c8f.js";import"./index.ba9b69da.js";import"./strings.ed5ba909.js";import"./isEqual.3dfeee36.js";import"./_Uint8Array.0caf8bf8.js";import"./_initCloneObject.e29c0e59.js";import"./drawer.f8ea4ac7.js";/* empty css             */import"./common.95eee561.js";import"./upload.55e3c344.js";import"./progress.d509d1cc.js";import"./_plugin-vue_export-helper.cdc0426e.js";import"./warningBar.3ea93790.js";import"./dropdown.70690357.js";import"./arrays.6dddb988.js";import"./index.c624c630.js";import"./date.34b11046.js";import"./dictionary.bf1d62cf.js";import"./dictionary.2f4c7964.js";import"./sysDictionary.20d54a34.js";const R={class:"gva-search-box"},F={class:"gva-table-box"},J={class:"gva-btn-list"},M=o("p",null,"确定要删除吗？",-1),G={style:{"text-align":"right","margin-top":"8px"}},H={class:"gva-pagination"},L=o("br",null,null,-1),N={class:"dialog-footer"},P={name:"Banner"},Q=Object.assign(P,{setup(P){const Q=e([]),W=e({image:"",path:"",sort:0,status:void 0}),X=a({image:[{required:!0,message:"",trigger:["input","blur"]}],status:[{required:!0,message:"",trigger:["input","blur"]}]}),Y=e(),Z=e(),$=e("https://huzhu.cnecip.com/api/");console.log($.value),console.log("https://huzhu.cnecip.com/api");const ee=()=>{Z.value.open()},ae=e=>{W.value.image=e},le=e(1),te=e(0),oe=e(10),se=e([]),re=e({}),ie=({prop:e,order:a})=>{re.value.sort=e,re.value.order=a,ce()},ue=()=>{re.value={},ce()},de=()=>{le.value=1,oe.value=10,ce()},ne=e=>{oe.value=e,ce()},pe=e=>{le.value=e,ce()},ce=async()=>{const e=await I({page:le.value,pageSize:oe.value,...re.value});0===e.code&&(se.value=e.data.list,te.value=e.data.total,le.value=e.data.page,oe.value=e.data.pageSize)};ce();(async()=>{Q.value=await O("status")})();const me=e([]),ve=e=>{me.value=e},fe=e(!1),ge=async()=>{const e=[];if(0===me.value.length)return void f({type:"warning",message:"请选择要删除的数据"});me.value&&me.value.map((a=>{e.push(a.ID)}));0===(await D({ids:e})).code&&(f({type:"success",message:"删除成功"}),se.value.length===e.length&&le.value>1&&le.value--,fe.value=!1,ce())},be=e(""),ye=async e=>{0===(await S({ID:e.ID})).code&&(f({type:"success",message:"删除成功"}),1===se.value.length&&le.value>1&&le.value--,ce())},he=e(!1),je=()=>{be.value="create",he.value=!0},_e=()=>{he.value=!1,W.value={image:"",path:"",sort:0,status:void 0}},we=async()=>{var e;null==(e=Y.value)||e.validate((async e=>{if(!e)return;let a;switch(be.value){case"create":default:a=await B(W.value);break;case"update":a=await q(W.value)}0===a.code&&(f({type:"success",message:"创建/更改成功"}),_e(),ce())}))};return(e,a)=>{const f=U,I=E,D=g,S=z,B=V,q=w,O=k,P=C,ce=_,Ce=b,ke=h,xe=j,Ve=y;return l(),t("div",null,[o("div",R,[s(S,{inline:!0,model:re.value,class:"demo-form-inline",onKeyup:u(de,["enter"])},{default:r((()=>[s(I,{label:"创建时间"},{default:r((()=>[s(f,{modelValue:re.value.startCreatedAt,"onUpdate:modelValue":a[0]||(a[0]=e=>re.value.startCreatedAt=e),type:"datetime",placeholder:"开始时间"},null,8,["modelValue"]),i(" — "),s(f,{modelValue:re.value.endCreatedAt,"onUpdate:modelValue":a[1]||(a[1]=e=>re.value.endCreatedAt=e),type:"datetime",placeholder:"结束时间"},null,8,["modelValue"])])),_:1}),s(I,null,{default:r((()=>[s(D,{type:"primary",icon:"search",onClick:de},{default:r((()=>[i("查询")])),_:1}),s(D,{icon:"refresh",onClick:ue},{default:r((()=>[i("重置")])),_:1})])),_:1})])),_:1},8,["model","onKeyup"])]),o("div",F,[o("div",J,[s(D,{type:"primary",icon:"plus",onClick:je},{default:r((()=>[i("新增")])),_:1}),s(B,{visible:fe.value,"onUpdate:visible":a[4]||(a[4]=e=>fe.value=e),placement:"top",width:"160"},{reference:r((()=>[s(D,{icon:"delete",style:{"margin-left":"10px"},disabled:!me.value.length,onClick:a[3]||(a[3]=e=>fe.value=!0)},{default:r((()=>[i("删除")])),_:1},8,["disabled"])])),default:r((()=>[M,o("div",G,[s(D,{type:"primary",link:"",onClick:a[2]||(a[2]=e=>fe.value=!1)},{default:r((()=>[i("取消")])),_:1}),s(D,{type:"primary",onClick:ge},{default:r((()=>[i("确定")])),_:1})])])),_:1},8,["visible"])]),s(P,{ref:"multipleTable",style:{width:"100%"},"tooltip-effect":"dark",data:se.value,"row-key":"ID",onSelectionChange:ve,onSortChange:ie},{default:r((()=>[s(q,{type:"selection",width:"55"}),s(q,{align:"left",label:"日期",width:"180"},{default:r((e=>[i(d(n(T)(e.row.CreatedAt)),1)])),_:1}),s(q,{align:"left",label:"图片",prop:"image",width:"180"},{default:r((e=>[s(O,{style:{width:"150px"},src:$.value+e.row.image},null,8,["src"])])),_:1}),s(q,{align:"left",label:"路径",prop:"path",width:"120"}),s(q,{sortable:"",align:"left",label:"排序",prop:"sort",width:"120"}),s(q,{align:"left",label:"状态",prop:"status",width:"120"},{default:r((e=>[i(d(n(K)(e.row.status,Q.value)),1)])),_:1}),s(q,{align:"left",label:"按钮组"},{default:r((e=>[s(D,{type:"primary",link:"",icon:"edit",class:"table-button",onClick:a=>(async e=>{const a=await A({ID:e.ID});be.value="update",0===a.code&&(W.value=a.data.rebanner,he.value=!0)})(e.row)},{default:r((()=>[i("变更")])),_:2},1032,["onClick"]),s(D,{type:"primary",link:"",icon:"delete",onClick:a=>{return l=e.row,void v.confirm("确定要删除吗?","提示",{confirmButtonText:"确定",cancelButtonText:"取消",type:"warning"}).then((()=>{ye(l)}));var l}},{default:r((()=>[i("删除")])),_:2},1032,["onClick"])])),_:1})])),_:1},8,["data"]),o("div",H,[s(ce,{layout:"total, sizes, prev, pager, next, jumper","current-page":le.value,"page-size":oe.value,"page-sizes":[10,30,50,100],total:te.value,onCurrentChange:pe,onSizeChange:ne},null,8,["current-page","page-size","total"])])]),s(Ve,{modelValue:he.value,"onUpdate:modelValue":a[8]||(a[8]=e=>he.value=e),"before-close":_e,title:"弹窗操作"},{footer:r((()=>[o("div",N,[s(D,{onClick:_e},{default:r((()=>[i("取 消")])),_:1}),s(D,{type:"primary",onClick:we},{default:r((()=>[i("确 定")])),_:1})])])),default:r((()=>[s(S,{model:W.value,"label-position":"right",ref_key:"elFormRef",ref:Y,rules:X,"label-width":"80px"},{default:r((()=>[s(I,{label:"图片:",prop:"image"},{default:r((()=>[s(O,{style:{width:"250px"},src:$.value+W.value.image},null,8,["src"]),L,s(D,{type:"primary",onClick:ee},{default:r((()=>[i("选择图片")])),_:1}),s(x,{ref_key:"chooseImgRef",ref:Z,onEnterImg:ae},null,512)])),_:1}),s(I,{label:"路径:",prop:"path"},{default:r((()=>[s(Ce,{modelValue:W.value.path,"onUpdate:modelValue":a[5]||(a[5]=e=>W.value.path=e),clearable:!0,placeholder:"请输入"},null,8,["modelValue"])])),_:1}),s(I,{label:"排序:",prop:"sort"},{default:r((()=>[s(Ce,{modelValue:W.value.sort,"onUpdate:modelValue":a[6]||(a[6]=e=>W.value.sort=e),modelModifiers:{number:!0},clearable:!0,placeholder:"请输入"},null,8,["modelValue"])])),_:1}),s(I,{label:"状态:",prop:"status"},{default:r((()=>[s(xe,{modelValue:W.value.status,"onUpdate:modelValue":a[7]||(a[7]=e=>W.value.status=e),placeholder:"请选择",style:{width:"100%"},clearable:!1},{default:r((()=>[(l(!0),t(p,null,c(Q.value,((e,a)=>(l(),m(ke,{key:a,label:e.label,value:e.value},null,8,["label","value"])))),128))])),_:1},8,["modelValue"])])),_:1})])),_:1},8,["model","rules"])])),_:1},8,["modelValue"])])}}});export{Q as default};
