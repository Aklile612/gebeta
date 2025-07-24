<script setup>
import { useForm, useFieldArray } from 'vee-validate'
import * as yup from 'yup'
import { ref } from 'vue'

// Schema
const schema = yup.object({
  title: yup.string().required('Recipe title is required'),
  description: yup.string().required('Description is required'),
  prepTime: yup.number().min(1, 'Must be at least 1 minute').required('Prep time is required'),
  cookTime: yup.number().min(1, 'Must be at least 1 minute').required('Cook time is required'),
  images: yup.array().min(1, 'At least one image is required'),
  category: yup.string().required('Category is required'),
  difficulty: yup.string().required('Difficulty is required'),
  ingredients: yup.array().of(
    yup.object({
      name: yup.string().required('Ingredient name is required'),
      quantity: yup.string().required('Quantity is required'),
    })
  ).min(1, 'At least one ingredient is required'),
  steps: yup.array().of(
    yup.object({
      description: yup.string().required('Step description is required'),
    })
  ).min(1, 'At least one step is required'),
  isPremium: yup.boolean(),
  price: yup.number().when('isPremium', {
    is: true,
    then: (schema) => schema.required('Price is required for premium recipes').min(1, 'Price must be at least 1'),
    otherwise: (schema) => schema.notRequired()
  })
})

// Form setup
const { handleSubmit, errors, values, setFieldValue, meta, touched } = useForm({
  validationSchema: schema,
  initialValues: {
    title: '',
    description: '',
    prepTime: '',
    cookTime: '',
    images: [],
    category: '',
    difficulty: '',
    ingredients: [{ name: '', quantity: '' }],
    steps: [{ description: '' }],
    isPremium: false,
    price: ''
  }
})

const { fields: ingredientFields, push: addIngredient, remove: removeIngredient } = useFieldArray('ingredients')
const { fields: stepFields, push: addStep, remove: removeStep } = useFieldArray('steps')

// File handling
function handleImageUpload(event) {
  const files = Array.from(event.target.files)
  setFieldValue('images', files)
}
const imagePreviews = ref([])

function handleImageUpload(event) {
  const files = Array.from(event.target.files)
  setFieldValue('images', files)

  // Generate preview URLs
  imagePreviews.value = files.map(file => ({
    url: URL.createObjectURL(file),
    name: file.name
  }))
}
// Submission
const onSubmit = handleSubmit((formValues) => {
  console.log('✅ Submitted:', formValues)
})

const removeImage = (index) => {
  const newImages = [...values.images]
  newImages.splice(index, 1)
  setFieldValue('images', newImages)
  imagePreviews.value.splice(index, 1)
}
</script>

<template>
  <div class="min-h-screen bg-[#fae3cd] py-8 px-4">
    <div class="max-w-3xl mx-auto">
      <!-- Header -->
      <div class="mb-6">
        <NuxtLink to="/" class="flex items-center text-gray-700 hover:text-orange-500 mb-2">
          <IconChevronLeft class="mr-1" />
          <span class="text-sm font-medium">Back</span>
        </NuxtLink>
        <h1 class="text-2xl font-bold text-gray-800">Create New Recipe</h1>
        <p class="text-gray-600">Share your culinary creation with the community</p>
      </div>

      <!-- Form -->
      <form @submit.prevent="onSubmit" class="bg-white rounded-lg shadow-md p-6">
        <!-- Basic Information Section -->
        <section class="mb-8">
          <h2 class="text-lg font-bold mb-4 pb-2 border-b border-gray-200">Basic Information</h2>
          
          <!-- Recipe Title -->
          <div class="mb-4">
            <label class="block text-sm font-medium text-gray-700 mb-1">Recipe Title*</label>
            <input 
              v-model="values.title" 
              placeholder="e.g., Grandma's Chocolate Chip Cookies" 
              class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-orange-500"
            />
            <span class="text-red-500 text-xs mt-1" v-if="touched?.title && errors.title">{{ errors.title }}</span>
          </div>
          
          <!-- Description -->
          <div class="mb-4">
            <label class="block text-sm font-medium text-gray-700 mb-1">Description*</label>
            <textarea 
              v-model="values.description" 
              placeholder="Tell us about your recipe..." 
              rows="3"
              class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-orange-500"
            ></textarea>
            <span class="text-red-500 text-xs mt-1" v-if="touched?.description && errors.description">{{ errors.description }}</span>
          </div>
          
          <!-- Category & Difficulty -->
          <div class="grid grid-cols-1 md:grid-cols-2 gap-4 mb-4">
            <div>
              <label class="block text-sm font-medium text-gray-700 mb-1">Category*</label>
              <select 
                v-model="values.category" 
                class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-orange-500"
              >
                <option value="">Select category</option>
                <option value="breakfast">Breakfast</option>
                <option value="lunch">Lunch</option>
                <option value="dinner">Dinner</option>
                <option value="dessert">Dessert</option>
                <option value="snack">Snack</option>
              </select>
              <span class="text-red-500 text-xs mt-1" v-if="touched?.category && errors.category">{{ errors.category }}</span>
            </div>
            
            <div>
              <label class="block text-sm font-medium text-gray-700 mb-1">Difficulty*</label>
              <select 
                v-model="values.difficulty" 
                class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-orange-500"
              >
                <option value="">Select difficulty</option>
                <option value="easy">Easy</option>
                <option value="medium">Medium</option>
                <option value="hard">Hard</option>
              </select>
              <span class="text-red-500 text-xs mt-1" v-if="touched?.difficulty && errors.difficulty">{{ errors.difficulty }}</span>
            </div>
          </div>
          
          <!-- Prep & Cook Time -->
          <div class="grid grid-cols-1 md:grid-cols-2 gap-4">
            <div>
              <label class="block text-sm font-medium text-gray-700 mb-1">Prep Time (minutes)*</label>
              <input 
                v-model="values.prepTime" 
                type="number" 
                min="1"
                class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-orange-500"
              />
              <span class="text-red-500 text-xs mt-1" v-if="touched?.prepTime && errors.prepTime">{{ errors.prepTime }}</span>
            </div>
            
            <div>
              <label class="block text-sm font-medium text-gray-700 mb-1">Cook Time (minutes)*</label>
              <input 
                v-model="values.cookTime" 
                type="number" 
                min="1"
                class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-orange-500"
              />
              <span class="text-red-500 text-xs mt-1" v-if="touched?.cookTime && errors.cookTime">{{ errors.cookTime }}</span>
            </div>
          </div>
        </section>

        <!-- Recipe Images Section -->
        <section class="mb-8">
          <h2 class="text-lg font-bold mb-4 pb-2 border-b border-gray-200">Recipe Images</h2>
          <div class="border-2 border-dashed border-gray-300 rounded-lg p-6 text-center">
            <div class="flex flex-col items-center justify-center">
              <svg class="w-12 h-12 text-gray-400 mb-2" fill="none" stroke="currentColor" viewBox="0 0 24 24" xmlns="http://www.w3.org/2000/svg">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 16l4.586-4.586a2 2 0 012.828 0L16 16m-2-2l1.586-1.586a2 2 0 012.828 0L20 14m-6-6h.01M6 20h12a2 2 0 002-2V6a2 2 0 00-2-2H6a2 2 0 00-2 2v12a2 2 0 002 2z"></path>
              </svg>
              <p class="text-sm text-gray-500 mb-2">Drag & drop images or click to browse</p>
              <p class="text-xs text-gray-400 mb-4">Support: JPG, PNG (Max 5MB each)</p>
              <input
                type="file"
                multiple
                accept="image/jpeg,image/png"
                @change="handleImageUpload"
                class="hidden"
                id="image-upload"
              />
              <label for="image-upload" class="px-4 py-2 bg-orange-500 text-white rounded-md cursor-pointer hover:bg-orange-600 transition">
                Choose Files
              </label>
            </div>
            <span class="text-red-500 text-xs mt-2" v-if="touched?.images && errors.images">{{ errors.images }}</span>
            <!-- Image Preview Thumbnails -->
            <div v-if="imagePreviews.length" class="mt-4 grid grid-cols-2 sm:grid-cols-3 gap-4">
              <div
                v-for="(image, index) in imagePreviews"
                :key="index"
                class="relative border rounded overflow-hidden"
              >
                <img :src="image.url" :alt="image.name" class="w-full h-32 object-cover" />
                <button
                  type="button"
                  @click="removeImage(index)"
                  class="absolute top-1 right-1 bg-white rounded-full p-1 shadow hover:bg-gray-100"
                >
                  <svg class="w-4 h-4 text-red-500" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12" />
                  </svg>
                </button>
              </div>
            </div>

          </div>
        </section>

        <!-- Ingredients Section -->
        <section class="mb-8">
          <h2 class="text-lg font-bold mb-4 pb-2 border-b border-gray-200">Ingredients</h2>
          <div class="space-y-3">
            <div v-for="(ingredient, idx) in ingredientFields" :key="ingredient.key" class="flex items-center gap-3">
              <input
                v-model="values.ingredients[idx].name"
                placeholder="Ingredient name"
                class="flex-1 px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-orange-500"
              />
              <input
                v-model="values.ingredients[idx].quantity"
                placeholder="Quantity"
                class="w-1/3 px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-orange-500"
              />
              <button
                type="button"
                @click="removeIngredient(idx)"
                class="text-red-500 hover:text-red-700"
              >
                <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24" xmlns="http://www.w3.org/2000/svg">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 7l-.867 12.142A2 2 0 0116.138 21H7.862a2 2 0 01-1.995-1.858L5 7m5 4v6m4-6v6m1-10V4a1 1 0 00-1-1h-4a1 1 0 00-1 1v3M4 7h16"></path>
                </svg>
              </button>
            </div>
            <button
              type="button"
              @click="addIngredient({ name: '', quantity: '' })"
              class="flex items-center text-orange-500 hover:text-orange-700 mt-2"
            >
              <svg class="w-5 h-5 mr-1" fill="none" stroke="currentColor" viewBox="0 0 24 24" xmlns="http://www.w3.org/2000/svg">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 6v6m0 0v6m0-6h6m-6 0H6"></path>
              </svg>
              Add ingredient
            </button>
            <span class="text-red-500 text-xs mt-1" v-if="touched?.ingredients && errors.ingredients">{{ errors.ingredients }}</span>
          </div>
        </section>

        <!-- Preparation Steps Section -->
        <section class="mb-8">
          <h2 class="text-lg font-bold mb-4 pb-2 border-b border-gray-200">Preparation Steps</h2>
          <div class="space-y-3">
            <div v-for="(step, idx) in stepFields" :key="step.key" class="flex gap-3">
              <div class="flex-1">
                <div class="flex items-start">
                  <span class="mt-2 mr-2 font-bold text-gray-500">{{ idx + 1 }}.</span>
                  <textarea
                    v-model="values.steps[idx].description"
                    placeholder="Describe this step..."
                    rows="2"
                    class="flex-1 px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-orange-500"
                  ></textarea>
                </div>
              </div>
              <button
                type="button"
                @click="removeStep(idx)"
                class="text-red-500 hover:text-red-700 mt-2"
              >
                <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24" xmlns="http://www.w3.org/2000/svg">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 7l-.867 12.142A2 2 0 0116.138 21H7.862a2 2 0 01-1.995-1.858L5 7m5 4v6m4-6v6m1-10V4a1 1 0 00-1-1h-4a1 1 0 00-1 1v3M4 7h16"></path>
                </svg>
              </button>
            </div>
            <button
              type="button"
              @click="addStep({ description: '' })"
              class="flex items-center text-orange-500 hover:text-orange-700 mt-2"
            >
              <svg class="w-5 h-5 mr-1" fill="none" stroke="currentColor" viewBox="0 0 24 24" xmlns="http://www.w3.org/2000/svg">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 6v6m0 0v6m0-6h6m-6 0H6"></path>
              </svg>
              Add Step
            </button>
            <span class="text-red-500 text-xs mt-1" v-if="touched?.steps && errors.steps">{{ errors.steps }}</span>
          </div>
        </section>

        <!-- Pricing Section -->
       
        <section class="mb-6">
          <h2 class="text-lg font-bold mb-4 pb-2 border-b border-gray-200">Pricing</h2>
          <div class="flex items-center mb-2">
            <input
              type="checkbox"
              id="premium-recipe"
              v-model="values.isPremium"
              class="w-4 h-4 text-orange-500 rounded focus:ring-orange-500"
            />
            <label for="premium-recipe" class="ml-2 text-sm font-medium text-gray-700">Premium Recipe</label>
          </div>
          <p class="text-xs text-gray-500 mb-3">Charge users to access this recipe</p>
          
          <!-- Price Input (Conditional) -->
          <div v-if="values.isPremium" class="mt-2 transition-all duration-200 ease-in-out">
            <label class="block text-sm font-medium text-gray-700 mb-1">Price (ETB)*</label>
            <div class="relative rounded-md shadow-sm">
              <div class="absolute inset-y-0 left-0 pl-3 flex items-center pointer-events-none">
                <span class="text-gray-500">ETB</span>
              </div>
              <input
                v-model="values.price"
                type="number"
                min="1"
                placeholder="Enter amount"
                class="block w-full pl-12 pr-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-orange-500"
              />
            </div>
            <span class="text-red-500 text-xs mt-1" v-if="touched?.price && errors.price">{{ errors.price }}</span>
          </div>
        </section>

        <!-- Submit Button -->
        <button
          type="submit"
          class="w-full bg-orange-500 text-white py-3 px-4 rounded-md hover:bg-orange-600 transition font-medium"
        >
          Submit Recipe
        </button>
      </form>
    </div>
  </div>
</template>

<style scoped>
/* Add any custom styles here if needed */
</style>