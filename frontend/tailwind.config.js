/** @type {import('tailwindcss').Config} */
export default {
  content: [
    "./index.html",
    "./src/**/*.{js,ts,jsx,tsx}",
  ],
  theme: {
    extend: {
      colors: {
        charade: {
          50: "#f5f6f9",
          100: "#e7e9f2",
          200: "#d5d7e8",
          300: "#b9bed7",
          400: "#969dc4",
          500: "#7d81b4",
          600: "#6b6ca5",
          700: "#625f96",
          800: "#54517c",
          900: "#464464",
          950: "#242331", // First Color
        },
        woodland: {
          50: "#f8f7ed",
          100: "#f0f0d7",
          200: "#e1e2b4",
          300: "#cdcf87",
          400: "#b6bb60",
          500: "#9a9f43",
          600: "#787e32",
          700: "#5b612a",
          800: "#474b24", // Second Color
          900: "#404423",
          950: "#21240f",
        },
        edward: {
          50: "#f7f8f8",
          100: "#f1f2f2",
          200: "#e3e7e7",
          300: "#cdd4d4",
          400: "#adb8b9",
          500: "#95a3a4", // Third Color
          600: "#798789",
          700: "#637173",
          800: "#545e60",
          900: "#485253",
          950: "#282e2f",
        },
        thunderbird: {
          50: "#fef3f2",
          100: "#fde6e3",
          200: "#fdd0cb",
          300: "#fab0a7",
          400: "#f58274",
          500: "#eb5948",
          600: "#d83d2a",
          700: "#c03221", // Fourth Color
          800: "#962b1e",
          900: "#7d291f",
          950: "#44110b",
        },
        "sandy-brown": {
          50: "#fef8ee",
          100: "#fdeed7",
          200: "#fbdaad",
          300: "#f8bf79",
          400: "#f49e4c", // Fifth Color
          500: "#f07d1f",
          600: "#e16315",
          700: "#bb4a13",
          800: "#953c17",
          900: "#783316",
          950: "#411709",
        },
        "athens-gray": {
          50: "#f6f7f8",
          100: "#e5e7eb", // Background Color
          200: "#dadde3",
          300: "#c1c7cf",
          400: "#a2aab8",
          500: "#8c93a5",
          600: "#7a8096",
          700: "#6e7287",
          800: "#5d5f70",
          900: "#4d4f5b",
          950: "#31323a",
        },
        "primary-color": "#271F30",
        "secondary-color": "#EEE3AB",
        "third-color": "#FF8600",
        "forth-color": "#934B00",
        "fifth-color": "#9FBBCC",
      },
    },
  },
  plugins: [],
}

