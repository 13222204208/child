/*! 
 Build based on gin-vue-admin 
 Time : 1680141068000 */
import{c1 as a,u as e,r as l,a as t,o as s,c as r,d as o,e as u,w as i,U as d,aC as m,a0 as p,h as c,j as n,b,E as f}from"./index.9fcd1cf4.js";import{E as v,a as j}from"./form-item.d732c3ca.js";/* empty css               *//* empty css              */import"./tag.591c2959.js";import{E as g,a as y}from"./select.838f2837.js";import"./scrollbar.aed9ef7e.js";/* empty css               */import{f as h,c as V,u as _}from"./banner.dbf5597e.js";import{g as w}from"./format.9eacff27.js";import"./_Uint8Array.660aa69b.js";import"./_initCloneObject.a1ac0bc9.js";import"./index.bb4181cc.js";import"./index.7c6e002b.js";import"./strings.bddfd219.js";import"./isEqual.06eb6a45.js";import"./date.34b11046.js";import"./dictionary.888f6128.js";import"./dictionary.3924a9ed.js";import"./sysDictionary.8f2d31b8.js";const x={class:"gva-form-box"},U={name:"Banner"},k=Object.assign(U,{setup(U){const k=a(),q=e(),C=l(""),E=l([]),D=l({image:"",path:"",sort:0,status:void 0}),O=t({image:[{required:!0,message:"",trigger:["input","blur"]}],status:[{required:!0,message:"",trigger:["input","blur"]}]}),A=l();(async()=>{if(k.query.id){const a=await h({ID:k.query.id});0===a.code&&(D.value=a.data.rebanner,C.value="update")}else C.value="create";E.value=await w("status")})();const B=async()=>{var a;null==(a=A.value)||a.validate((async a=>{if(!a)return;let e;switch(C.value){case"create":default:e=await V(D.value);break;case"update":e=await _(D.value)}0===e.code&&n({type:"success",message:"创建/更改成功"})}))},F=()=>{q.go(-1)};return(a,e)=>{const l=b,t=v,n=g,h=y,V=f,_=j;return s(),r("div",null,[o("div",x,[u(_,{model:D.value,ref_key:"elFormRef",ref:A,"label-position":"right",rules:O,"label-width":"80px"},{default:i((()=>[u(t,{label:"图片:",prop:"image"},{default:i((()=>[u(l,{modelValue:D.value.image,"onUpdate:modelValue":e[0]||(e[0]=a=>D.value.image=a),clearable:!1,placeholder:"请输入"},null,8,["modelValue"])])),_:1}),u(t,{label:"路径:",prop:"path"},{default:i((()=>[u(l,{modelValue:D.value.path,"onUpdate:modelValue":e[1]||(e[1]=a=>D.value.path=a),clearable:!0,placeholder:"请输入"},null,8,["modelValue"])])),_:1}),u(t,{label:"排序:",prop:"sort"},{default:i((()=>[u(l,{modelValue:D.value.sort,"onUpdate:modelValue":e[2]||(e[2]=a=>D.value.sort=a),modelModifiers:{number:!0},clearable:!0,placeholder:"请输入"},null,8,["modelValue"])])),_:1}),u(t,{label:"状态:",prop:"status"},{default:i((()=>[u(h,{modelValue:D.value.status,"onUpdate:modelValue":e[3]||(e[3]=a=>D.value.status=a),placeholder:"请选择",clearable:!1},{default:i((()=>[(s(!0),r(d,null,m(E.value,((a,e)=>(s(),p(n,{key:e,label:a.label,value:a.value},null,8,["label","value"])))),128))])),_:1},8,["modelValue"])])),_:1}),u(t,null,{default:i((()=>[u(V,{type:"primary",onClick:B},{default:i((()=>[c("保存")])),_:1}),u(V,{type:"primary",onClick:F},{default:i((()=>[c("返回")])),_:1})])),_:1})])),_:1},8,["model","rules"])])])}}});export{k as default};