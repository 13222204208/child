/*! 
 Build based on gin-vue-admin 
 Time : 1679909256000 */
System.register(["./menuItem-legacy.3323476d.js","./asyncSubmenu-legacy.f4938e31.js","./index-legacy.e17276f9.js","./tooltip-legacy.43e4e549.js","./popper-legacy.85337370.js","./icon-legacy.b28fcc2b.js","./_plugin-vue_export-helper-legacy.649209c3.js","./index-legacy.e2dcb181.js","./index-legacy.eda54a47.js","./index-legacy.05d259e6.js"],(function(e,n){"use strict";var t,o,r,u,c,i,f,l,a,s,d,p,y;return{setters:[function(e){t=e.default},function(e){o=e.default},function(e){r=e.z,u=e.aA,c=e.o,i=e.a0,f=e.w,l=e.c,a=e.U,s=e.aC,d=e.f,p=e.a1,y=e.J},function(){},function(){},function(){},function(){},function(){},function(){},function(){}],execute:function(){var n={name:"AsideComponent"};e("default",Object.assign(n,{props:{routerInfo:{type:Object,default:function(){return null}},isCollapse:{default:function(){return!1},type:Boolean},theme:{default:function(){return{}},type:Object}},setup:function(e){var n=e,h=r((function(){return n.routerInfo.children&&n.routerInfo.children.filter((function(e){return!e.hidden})).length?o:t}));return function(n,t){var o=u("AsideComponent");return e.routerInfo.hidden?d("",!0):(c(),i(p(y(h)),{key:0,"is-collapse":e.isCollapse,theme:e.theme,"router-info":e.routerInfo},{default:f((function(){return[e.routerInfo.children&&e.routerInfo.children.length?(c(!0),l(a,{key:0},s(e.routerInfo.children,(function(n){return c(),i(o,{key:n.name,"is-collapse":!1,"router-info":n,theme:e.theme},null,8,["router-info","theme"])})),128)):d("",!0)]})),_:1},8,["is-collapse","theme","router-info"]))}}}))}}}));