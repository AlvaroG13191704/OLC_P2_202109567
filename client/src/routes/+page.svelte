<script>
// @ts-nocheck

	import ActionBar from "./components/ActionBar.svelte";
	import Console from "./components/Console.svelte";
	import Editor from "./components/Editor.svelte";
	import Navbar from "./components/Navbar.svelte";
  import { dataContext } from "../store";
	import EditorCompiled from "./components/EditorCompiled.svelte";
</script>

<div class="flex w-full flex-col ">
  <Navbar />
  <div class="flex flex-row">
    <div class="flex flex-col grow">
      <Editor />
      <!-- Add the three  -->
      <Console />
    </div>
    <!-- Add the text area to the compile section -->
    <div class="bg-gray-900">
      <textarea class="bg-gray-600 text-gray-300 h-[650px] w-[600px]  p-3 font-bold text-lg" disabled bind:value={$dataContext.outputCompile}  />
      <!-- <div class="w-[600px]">

        <EditorCompiled bind:responseCodeCompiled={$dataContext.outputCompile} />
      </div> -->
      <h1 class=" text-gray-200 px-2 py-1 text-2xl rounded-m">Código compilado</h1>
    </div>
    <div class="flex flex-col grow w-[250px] ">
      <ActionBar />
    </div>
  </div>
</div>