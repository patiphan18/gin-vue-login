<template>
  <div v-if="!auth.user" class="flex flex-col items-center justify-center gap-4 w-96 bg-white p-8 rounded-md shadow-md">
    <h2 class="text-2xl font-bold">Login</h2>
    <p v-if="error" class="text-red-500">{{ error }}</p>
    <input v-model="username" placeholder="username" class="border border-gray-300 rounded-md p-2 w-full" />
    <input v-model="password" type="password" placeholder="password" class="border border-gray-300 rounded-md p-2 w-full" />
    <button @click="handleLogin" class="bg-blue-500 text-white rounded-md p-2 w-full flex items-center justify-center" :disabled="isLoading">
      <span v-if="isLoading">
        <svg class="animate-spin h-5 w-5 mr-3 ..." viewBox="0 0 24 24">
          <circle class="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" stroke-width="4"></circle>
          <path class="opacity-75" fill="currentColor" d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4zm2 5.291A7.962 7.962 0 014 12H0c0 3.042 1.135 5.824 3 7.938l3-2.647z"></path>
        </svg>
      </span>
      <span v-else>ลงชื่อเข้าใช้งาน</span>
    </button>
    <a href="/register" class="text-sm text-blue-500">สมัครสมาชิก</a>
  </div>
</template>

<script setup>
import { onMounted, ref } from "vue";
import { useAuthStore } from "../stores/auth";
import { useRouter } from "vue-router";

const username = ref("");
const password = ref("");
const isLoading = ref(false);
const error = ref(null);

const auth = useAuthStore();
const router = useRouter();

onMounted(() => {
  if (auth.user) {
    router.push("/profile");
  }
});

const handleLogin = async () => {
  isLoading.value = true;
  error.value = "";
  try {
    await auth.login(username.value, password.value);
    router.push("/profile");
  } catch (err) {
    error.value = err.response.data.error;
  } finally {
    isLoading.value = false;
  }
};
</script>