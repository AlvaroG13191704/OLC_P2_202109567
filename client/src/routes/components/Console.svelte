
<script>
	import { dataContext } from "../../store";
	import ErrorConsole from "./consoleComponents/ErrorConsole.svelte";
	import SymbolConsole from "./consoleComponents/SymbolConsole.svelte";



  let activeConsole = 'main' 


</script>



<div class="bg-gray-900 min-h-screen text-lg p-2 "> 
  <div class="m-2">
    <button class="bg-gray-700 text-gray-200 px-2 py-1 rounded-md hover:bg-gray-500" on:click={() => activeConsole = 'main'}>Consola</button>
    <button class="bg-gray-700 text-gray-200 px-2 py-1 rounded-md hover:bg-gray-500" on:click={() => activeConsole = 'symbols'}>Simbolos</button>
    <button class="bg-gray-700 text-gray-200 px-2 py-1 rounded-md hover:bg-gray-500" on:click={() => activeConsole = 'errors'}>Errores</button>
  </div>
  
  <div class="border-b-2 border-gray-700  m-2">
  {#if activeConsole === 'main'}
    <textarea class="bg-slate-300 text-black w-full  p-3 font-bold" disabled  />
  {:else if activeConsole === 'symbols'}
    <SymbolConsole  />
  {:else if activeConsole === 'errors'}
    <ErrorConsole  />
  {/if}        
  </div>
  

</div>