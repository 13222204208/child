/*! 
 Build based on gin-vue-admin 
 Time : 1679909256000 */
import{q as e,s as t,A as s,C as a,z as r,o as i,c as o,n,J as l,I as d,f as c,a2 as p,_ as u,a5 as v}from"./index.61dd0a2d.js";const f=e({direction:{type:String,values:["horizontal","vertical"],default:"horizontal"},contentPosition:{type:String,values:["left","center","right"],default:"center"},borderStyle:{type:t(String),default:"solid"}}),m=s({name:"ElDivider"});const y=v(u(s({...m,props:f,setup(e){const t=e,s=a("divider"),u=r((()=>s.cssVar({"border-style":t.borderStyle})));return(e,t)=>(i(),o("div",{class:n([l(s).b(),l(s).m(e.direction)]),style:p(l(u)),role:"separator"},[e.$slots.default&&"vertical"!==e.direction?(i(),o("div",{key:0,class:n([l(s).e("text"),l(s).is(e.contentPosition)])},[d(e.$slots,"default")],2)):c("v-if",!0)],6))}}),[["__file","/home/runner/work/element-plus/element-plus/packages/components/divider/src/divider.vue"]]));export{y as E};
