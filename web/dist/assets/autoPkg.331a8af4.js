/*! 
 Build based on gin-vue-admin 
 Time : 1679909256000 */
import{r as e,o as a,c as l,e as t,d as o,w as s,h as i,j as r,bm as d,E as p,b as c}from"./index.61dd0a2d.js";import{E as u}from"./dialog.a2b2142c.js";import"./overlay.533fa516.js";import{E as m,a as n}from"./form-item.a3ced827.js";/* empty css              */import{E as f,a as b}from"./table-column.720e0d1a.js";import"./checkbox.384d52e9.js";import"./tag.3ecd1fed.js";/* empty css                *//* empty css               */import"./scrollbar.ef9dc6d6.js";/* empty css               */import{i as g,d as v,j}from"./autoCode.3f49bdae.js";import{W as k}from"./warningBar.3ea93790.js";import{_ as w}from"./_plugin-vue_export-helper.cdc0426e.js";import"./refs.1f88eb22.js";import"./_Uint8Array.0caf8bf8.js";import"./_initCloneObject.e29c0e59.js";import"./index.cbcb4c8f.js";import"./isEqual.3dfeee36.js";/* empty css             */const y={class:"gva-table-box"},_={class:"gva-btn-list"},h={class:"dialog-footer"},V={name:"AutoPkg"},x=w(Object.assign(V,{setup(w){const V=e({packageName:"",label:"",desc:""}),x=e({packageName:[{required:!0,message:"请输入包名",trigger:"blur"},{validator:(e,a,l)=>{/^\d+$/.test(a[0])?l(new Error("不能够以数字开头")):l()},trigger:"blur"}]}),C=e(!1),N=()=>{C.value=!1,V.value={packageName:"",label:"",desc:""}},E=e(null),B=async()=>{E.value.validate((async e=>{if(e){0===(await g(V.value)).code&&r({type:"success",message:"添加成功",showClose:!0}),T(),N()}}))},U=e([]),T=async()=>{const e=await v();0===e.code&&(U.value=e.data.pkgs)};return T(),(e,g)=>{const v=p,w=f,P=b,q=c,A=m,I=n,O=u;return a(),l("div",null,[t(k,{href:"https://www.bilibili.com/video/BV1kv4y1g7nT?p=3",title:"此功能为开发环境使用，不建议发布到生产，具体使用效果请看视频https://www.bilibili.com/video/BV1kv4y1g7nT?p=3"}),o("div",y,[o("div",_,[t(v,{type:"primary",icon:"plus",onClick:g[0]||(g[0]=e=>{C.value=!0})},{default:s((()=>[i("新增")])),_:1})]),t(P,{data:U.value},{default:s((()=>[t(w,{align:"left",label:"id",width:"60",prop:"ID"}),t(w,{align:"left",label:"包名",width:"150",prop:"packageName"}),t(w,{align:"left",label:"展示名",width:"150",prop:"label"}),t(w,{align:"left",label:"描述","min-width":"150",prop:"desc"}),t(w,{align:"left",label:"操作",width:"200"},{default:s((e=>[t(v,{icon:"delete",type:"primary",link:"",onClick:a=>(async e=>{d.confirm("此操作仅删除数据库中的pkg存储，后端相应目录结构请自行删除与数据库保持一致！","提示",{confirmButtonText:"确定",cancelButtonText:"取消",type:"warning"}).then((async()=>{0===(await j(e)).code&&(r({type:"success",message:"删除成功!"}),T())}))})(e.row)},{default:s((()=>[i("删除")])),_:2},1032,["onClick"])])),_:1})])),_:1},8,["data"])]),t(O,{modelValue:C.value,"onUpdate:modelValue":g[4]||(g[4]=e=>C.value=e),"before-close":N,title:"创建Package"},{footer:s((()=>[o("div",h,[t(v,{onClick:N},{default:s((()=>[i("取 消")])),_:1}),t(v,{type:"primary",onClick:B},{default:s((()=>[i("确 定")])),_:1})])])),default:s((()=>[t(k,{title:"新增Pkg用于自动化代码使用"}),t(I,{ref_key:"pkgForm",ref:E,model:V.value,rules:x.value,"label-width":"80px"},{default:s((()=>[t(A,{label:"包名",prop:"packageName"},{default:s((()=>[t(q,{modelValue:V.value.packageName,"onUpdate:modelValue":g[1]||(g[1]=e=>V.value.packageName=e),autocomplete:"off"},null,8,["modelValue"])])),_:1}),t(A,{label:"展示名",prop:"label"},{default:s((()=>[t(q,{modelValue:V.value.label,"onUpdate:modelValue":g[2]||(g[2]=e=>V.value.label=e),autocomplete:"off"},null,8,["modelValue"])])),_:1}),t(A,{label:"描述",prop:"desc"},{default:s((()=>[t(q,{modelValue:V.value.desc,"onUpdate:modelValue":g[3]||(g[3]=e=>V.value.desc=e),autocomplete:"off"},null,8,["modelValue"])])),_:1})])),_:1},8,["model","rules"])])),_:1},8,["modelValue"])])}}}),[["__scopeId","data-v-1a433cdb"]]);export{x as default};
