<script setup>
  definePageMeta({
    layout:false
  })
</script>


<template>
  <div class="flex min-h-screen justify-center">
    <!-- image background -->
    <div class="hero w-1/2 min-h-screen bg-no-repeat bg-center bg-cover" style="background-image: url('/image/register.png');">
      <div class="flex flex-col justify-end items-center text-neutral-content w-full h-full pb-10 text-center">
        <div class="max-w-md">
          <p class="mb-5 font-bold text-2xl">Every Recipe Has a Story <div class="border-t-2 border-white w-2"> <br> Share yours</div>Share </p>
          <p class="mb-5  text-lg"></p>
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
      <p class="text-5xl font-bold">Create Your Account</p>
      <p class="font-bold text-sm">Start Your Curlinary Journey today</p>
    </div>

    <!-- Form -->
    <form @submit.prevent="handleLogin" class="w-full">
      <fieldset class="bg-[#fae3cd]   w-[35vw] p-6 mx-5">
        <!-- Removed fieldset style classes that gave dark bg -->
        <label class="label text-lg" for="email">Email</label>
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
          class=" bg-orange-500 hover:bg-orange-600 text-white flex justify-center items-center group hover:scale-110 ease-out transition-all w-[300px] h-12  text-center  text-sm  font-semibold rounded-[6px]"
        >
          SignUp
        </button>
      </fieldset>
    </form>

    <!-- Link to Register -->
    <div>
      <NuxtLink to="/login">
        Already Have an account
        <span class="text-orange-500 hover:underline">Log In</span>
      </NuxtLink>
    </div>
  </div>
</div>
  </div>
  
  </template>
  