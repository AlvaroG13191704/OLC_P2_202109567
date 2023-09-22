<script>
	import ActionBar from "./components/ActionBar.svelte";
	import Console from "./components/Console.svelte";
	import Editor from "./components/Editor.svelte";
	import Navbar from "./components/Navbar.svelte";

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
      <textarea class="bg-gray-600 text-black h-[650px] w-[500px]  p-3 font-bold" disabled />
      <h1 class=" text-gray-200 px-2 py-1 text-2xl rounded-m">Código compilado</h1>
    </div>
    <div class="flex flex-col grow w-[400px] ">
      <ActionBar />
    </div>
  </div>
</div>