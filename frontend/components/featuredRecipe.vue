<script setup>
    defineProps({
        title: String,
        image: String,
        diffculty: String,
        cookTime: Number,
        prepTime:Number,
        rating: Number,
    })

    const isLiked = ref(false)
    const isBookmarked = ref(false)

    function toggleLike() {
      isLiked.value = !isLiked.value
    }

    function toggleBookmark() {
      isBookmarked.value = !isBookmarked.value
    }
</script>
<!-- #FFF6ED -->

<template>
    <div class="w-full h-[97vh] bg-[#fae3cd]">
        <div class="mx-[30vw] pt-[40px]">
            <h1 class="mb-5 text-4xl ml-[7vw] font-serif  font-bold">Featured Recipes</h1>
            <p class="text-lg font-semibold ml-5 text-gray-500 -mt-2">Discover the most loved recipe from the community</p>
        </div>
        <div class="ml-10 flex gap-10 justify-center  mt-12">
            <div class="card card-md bg-white w-80 h-[55vh] shadow-sm rounded-xl overflow-hidden">
                <!-- Image -->
                <figure>
                  <img
                    class="w-full h-[180px] object-cover hover:scale-110  transition-all rounded-lg"
                    src="/image/hero.png"
                    alt="food image"
                  />
                </figure>
            
                <!-- Body -->
                <div class="card-body p-4 pt-2">
                  <!-- like and bookmark -->
                  <div class="flex justify-end gap-2 mb-2">
                    <!-- Like button -->
                    <button @click="toggleLike">
                      <IconHeart :class="[isLiked ? 'text-red-500' : '']" class="w-5 h-5 hover:scale-110 transition-all" />
                    </button>

                    <!-- Bookmark button -->
                    <button @click="toggleBookmark">
                      <IconBookmark :class="[isBookmarked ? 'text-blue-500' : '']" class="w-5 h-5 hover:scale-110 transition-all" />
                    </button>
                  </div>
                  <!-- Top Row: Rating -->
                  <div class="flex items-center gap-1 text-sm text-yellow-500">
                    <IconStar class="w-4 h-4" />
                    <span class="text-xs font-medium">4.5</span>
                  </div>
              
                  <!-- Title + Badge -->
                  <div class="flex items-center justify-between mt-1">
                    <h2 class="text-md font-bold text-gray-800">Creamy Truffle Pasta</h2>
                    <span class="badge badge-warning text-[10px]">Premium</span>
                  </div>
              
                  <!-- Description -->
                  <p class="text-sm text-gray-600 mt-1">
                    A rich and indulgent pasta dish with truffle cream sauce
                  </p>
              
                  <!-- Footer Row: User + Time -->
                  <div class="flex items-center justify-between mt-1 text-xs text-gray-500">
                    <div class="flex items-center gap-1">
                        <IconUsers class="w-4 h-4 text-gray-500" />
                      <span>By User</span>
                    </div>
                    <div class="flex items-center gap-1">
                      <IconClock9 class="w-4 h-4" />
                      <span>29 min</span>
                    </div>
                  </div>
                </div>
            </div>

            <div class="card card-md bg-white w-80 h-[55vh] shadow-sm rounded-xl overflow-hidden">
                <!-- Image -->
                <figure>
                  <img
                    class="w-full h-[180px] object-cover hover:scale-110  transition-all rounded-lg"
                    src="/image/cake1.png"
                    alt="food image"
                  />
                </figure>
            
                <!-- Body -->
                <div class="card-body p-4 pt-2">
                
                  <!-- Top Row: Rating -->
                  <div class="flex items-center gap-1 text-sm text-yellow-500">
                    <IconStar class="w-4 h-4" />
                    <span class="text-xs font-medium">4.5</span>
                  </div>
              
                  <!-- Title + Badge -->
                  <div class="flex items-center justify-between mt-1">
                    <h2 class="text-md font-bold text-gray-800">chocklet berry cake</h2>
                  </div>
              
                  <!-- Description -->
                  <p class="text-sm text-gray-600 mt-1">
                    Decdent berry cake toked with berry cakes
                  </p>
              
                  <!-- Footer Row: User + Time -->
                  <div class="flex items-center justify-between mt-1 text-xs text-gray-500">
                    <div class="flex items-center gap-1">
                        <IconUsers class="w-4 h-4 text-gray-500" />
                        <span>By User</span>
                    </div>
                    <div class="flex items-center gap-1">
                      <IconClock9 class="w-4 h-4" />
                      <span>29 min</span>
                    </div>
                  </div>
                </div>
            </div>
            <div class="card card-md bg-white w-80 h-[55vh] shadow-sm rounded-xl overflow-hidden">
                <!-- Image -->
                <figure>
                  <img
                    class="w-full h-[180px] object-cover hover:scale-110  transition-all rounded-lg"
                    src="/image/vegs.png"
                    alt="food image"
                  />
                </figure>
            
                <!-- Body -->
                <div class="card-body p-4 pt-2">
                
                  <!-- Top Row: Rating -->
                  <div class="flex items-center gap-1 text-sm text-yellow-500">
                    <IconStar class="w-4 h-4" />
                    <span class="text-xs font-medium">4.5</span>
                  </div>
              
                  <!-- Title + Badge -->
                  <div class="flex items-center justify-between mt-1">
                    <h2 class="text-md font-bold text-gray-800">Rainbow Buddha Bowl</h2>
                  </div>
              
                  <!-- Description -->
                  <p class="text-sm text-gray-600 mt-1">
                    Nutrstios and colorfull bowl packed with fresh ingredients I hope it works
                  </p>
              
                  <!-- Footer Row: User + Time -->
                  <div class="flex items-center justify-between mt-1 text-xs text-gray-500">
                    <div class="flex items-center gap-1">
                        <IconUsers class="w-4 h-4 text-gray-500" />
                        <span>By User</span>
                    </div>
                    <div class="flex items-center gap-1">
                      <IconClock9 class="w-4 h-4" />
                      <span>29 min</span>
                    </div>
                  </div>
                </div>
            </div>
            
        </div>
    </div>
</template>