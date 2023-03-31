/*! 
 Build based on gin-vue-admin 
 Time : 1680141068000 */
import{L as e}from"./index.9fcd1cf4.js";const t=100,o=600,n={beforeMount(n,a){const d=a.value,{interval:r=t,delay:s=o}=e(d)?{}:d;let c,i;const l=()=>e(d)?d():d.handler(),u=()=>{i&&(clearTimeout(i),i=void 0),c&&(clearInterval(c),c=void 0)};n.addEventListener("mousedown",(e=>{0===e.button&&(u(),l(),document.addEventListener("mouseup",(()=>u()),{once:!0}),i=setTimeout((()=>{c=setInterval((()=>{l()}),r)}),s))}))}};export{n as v};
