import create from "zustand";
const useStore = create((set) => ({
  cari: "",
  setCari: (cari) => set((state) => ({ ...state, cari })),
}));

export default useStore;
