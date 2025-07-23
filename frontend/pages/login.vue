<script setup>
  definePageMeta({
    layout:false
  })
</script>


<template>
  <div class="flex min-h-screen justify-center">
    <!-- image background -->
    <div class="hero w-1/2 min-h-screen bg-no-repeat bg-center bg-cover" style="background-image: url('/image/login.png');">
      <div class="flex flex-col justify-end items-center text-neutral-content w-full h-full pb-10 text-center">
        <div class="max-w-md">
          <p class="mt-7 mr-4 font-bold text-2xl">Discover Amazing Recipes</p>
          <p class="mt-7 mr-4 text-lg">Join thousands of food lovers sharing their favorite recipes</p>
        </div>
      </div>
    </div>

  
    <!-- Optional Other Side -->
        <!-- Right Side (Form) -->
        <div class="w-1/2 bg-[#fae3cd] flex items-center justify-center">
  <div class="flex flex-col items-center space-y-6 -mt-6"> <!-- reduced vertical space with -mt -->
    <!-- Icon + Gebeta Name -->
    <div class="flex flex-col items-center">
      <div class="w-16 h-16 bg-orange-500 rounded-full flex items-center justify-center shadow-md">
        <IconUtensils class="text-white w-8 h-8" />
      </div>
      <p class="text-2xl mt-2 font-mono font-bold">Gebeta</p>
      <p class="text-sm text-gray-600">Recipe Sharing Community</p>
    </div>

    <!-- Welcome Text -->
    <div class="text-center">
      <p class="text-3xl font-bold">Welcome Back</p>
      <p class="font-bold text-sm">Login to share your next delicious recipe!</p>
    </div>

    <!-- Form -->
    <form @submit.prevent="handleLogin" class="w-full">
      <fieldset class="bg-[#fae3cd]   w-[40vw] p-6 ">
        <!-- Removed fieldset style classes that gave dark bg -->
        <label class="label text-lg" for="email">Enter Your Email</label>
        <div class="relative mb-4">
          <input
            id="email"
            type="email"
            class="input input-bordered bg-white w-full pl-10"
            placeholder="Email"
            v-model="email"
          />
          <IconMail class="absolute left-3 top-1/2 transform -translate-y-1/2 text-gray-500" />
        </div>

        <label class="label text-lg" for="password">Password</label>
        <div class="relative mb-4">
          <input
            id="password"
            type="password"
            class="input input-bordered bg-white w-full pl-10"
            placeholder="Password"
            v-model="password"
          />
          <IconLock class="absolute left-3 top-1/2 transform -translate-y-1/2 text-gray-500" />
        </div>

        <button
          type="submit"
          class="btn bg-orange-500 hover:bg-orange-600 text-white w-full mt-2"
        >
          Login
        </button>
      </fieldset>
    </form>

    <!-- Link to Register -->
    <div>
      <NuxtLink to="/register">
        <span class="font-semibold ">New Here?</span>
        <span class="text-range-500 font-semibold hover:underline">Create an account</span>
      </NuxtLink>
    </div>
  </div>
</div>
  </div>
  
  </template>
  