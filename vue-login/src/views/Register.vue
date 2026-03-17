<template>
  <div v-if="!auth.user" class="flex flex-col items-center justify-center gap-4 w-96 bg-white p-8 rounded-md shadow-md">
    <h2 class="text-2xl text-center font-bold">Register</h2>
    <p v-if="error" class="text-red-500">{{ error }}</p>
    <input v-model="username" placeholder="username" class="border border-gray-300 rounded-md p-2 w-full" />
    <input v-model="password" type="password" placeholder="password" class="border border-gray-300 rounded-md p-2 w-full" />
    <input v-model="confirmPassword" type="password" placeholder="confirm password" class="border border-gray-300 rounded-md p-2 w-full" />
    <button @click="handleRegister" class="bg-blue-500 text-white rounded-md p-2 w-full flex items-center justify-center" :disabled="isLoading">
      <span v-if="isLoading">
        <svg class="animate-spin h-5 w-5 mr-3 ..." viewBox="0 0 24 24">
          <circle class="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" stroke-width="4"></circle>
          <path class="opacity-75" fill="currentColor" d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4zm2 5.291A7.962 7.962 0 014 12H0c0 3.042 1.135 5.824 3 7.938l3-2.647z"></path>
        </svg>
      </span>
      <span v-else>สมัครสมาชิก</span>
    </button>
    <a href="/login" class="text-sm text-blue-500">เข้าสู่ระบบ</a>
  </div>
</template>

<script setup>
import { ref, onMounted } from "vue";
import { useAuthStore } from "../stores/auth";
import { useRouter } from "vue-router";

const username = ref("");
const password = ref("");
const confirmPassword = ref("");
const error = ref(null);
const isLoading = ref(false);

const auth = useAuthStore();
const router = useRouter();

onMounted(() => {
  if (auth.user) {
    router.push("/profile");
  }
});

const handleRegister = async () => {
  isLoading.value = true;
  error.value = "";
  try {
    await auth.register(username.value, password.value, confirmPassword.value);
    router.push("/login");
  } catch (err) {
    error.value = err.response.data.error;
  } finally {
    isLoading.value = false;
  }
};
</script>
