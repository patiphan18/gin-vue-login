<template>
  <div v-if="auth.user" class="flex flex-col items-center justify-center gap-8 w-96 bg-white p-8 rounded-md shadow-md">
    <div class="text-lg text-gray-500">Welcome User: {{ auth.user.username }}</div>
    <button @click="logout" class="bg-blue-500 text-white rounded-md p-2 w-full flex items-center justify-center">
      ออกจากระบบ
    </button>
  </div>
</template>

<script setup>
import { onMounted, ref } from "vue";
import { useAuthStore } from "../stores/auth";
import { useRouter } from "vue-router";

const auth = useAuthStore();
const router = useRouter();

onMounted(async () => {
  try {
    await auth.fetchProfile();
  } catch (err) {
    router.push("/login");
  }
});

const logout = () => {
  auth.logout();
  router.push("/login");
};
</script>