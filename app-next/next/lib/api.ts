import axios from "axios";

export const Axios = axios.create({
  // baseURL: "http://localhost:3000",
  baseURL: process.env.NEXT_BASE_URL,
  responseType: "json",
  headers: {
    "Content-Type": "application/json",
  },
});