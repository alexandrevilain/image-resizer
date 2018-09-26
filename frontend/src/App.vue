<template>
  <div id="app">
    <notifications group="all" position='bottom left'/>
    <div class="d-flex flex-column flex-md-row align-items-center p-3 px-md-4 mb-3 bg-white border-bottom box-shadow">
      <h5 class="my-0 mr-md-auto font-weight-normal">Image resizer</h5>
      <nav class="my-2 my-md-0 mr-md-3">
        <router-link to="/" class="p-2 text-dark">Home</router-link>
        <router-link to="/convert" class="p-2 text-dark">Resize an image</router-link>
        <router-link to="/finished" class="p-2 text-dark">Resized images</router-link>
      </nav>
    </div>
    <div class="container">
      <div class="row">
        <router-view/>
      </div>
    </div>
  </div>
</template>

<style>
</style>
