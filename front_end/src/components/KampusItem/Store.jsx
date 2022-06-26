import create from "zustand";
import axios from "axios";

const baseUrl = "https://reqres.in";
const token = localStorage.getItem("token");

const useStore = create((set) => ({
  cari: "",
  setCari: (cari) => set((state) => ({ ...state, cari })),

  token: localStorage.getItem("token"),
  data: [],
  getdata: async () => {
    const res = await axios.get(`${baseUrl}/api/users`, {
      headers: {
        Authorization: token,
      },
    });
    set((state) => ({ ...state, data: res.data.data }));
  },
}));
export default useStore;
