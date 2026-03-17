import { defineStore } from "pinia";
import api from "../api/axios";

export const useAuthStore = defineStore("auth", {
  state: () => ({
    user: null,
    token: localStorage.getItem("token") || null,
  }),

  actions: {
    async login(username: string, password: string) {
      const res = await api.post("/login", {
        username,
        password,
      });

      this.token = res.data.token;
      if (this.token) {
        localStorage.setItem("token", this.token);
      }
    },

    async register(username: string, password: string, confirmPassword: string) {
      await api.post("/register", {
        username,
        password,
        confirmPassword,
      });
    },

    async fetchProfile() {
      const res = await api.get("/api/profile");
      this.user = res.data;
    },

    logout() {
      this.token = null;
      this.user = null;
      localStorage.removeItem("token");
    },
  },
});