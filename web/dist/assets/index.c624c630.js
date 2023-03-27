/*! 
 Build based on gin-vue-admin 
 Time : 1679909256000 */
import{L as e}from"./index.61dd0a2d.js";const t=100,o=600,n={beforeMount(n,d){const a=d.value,{interval:r=t,delay:s=o}=e(a)?{}:a;let i,l;const u=()=>e(a)?a():a.handler(),v=()=>{l&&(clearTimeout(l),l=void 0),i&&(clearInterval(i),i=void 0)};n.addEventListener("mousedown",(e=>{0===e.button&&(v(),u(),document.addEventListener("mouseup",(()=>v()),{once:!0}),l=setTimeout((()=>{i=setInterval((()=>{u()}),r)}),s))}))}};export{n as v};
