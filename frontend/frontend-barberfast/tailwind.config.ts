import type { Config } from "tailwindcss";

export default {
  content: [
    "./src/pages/**/*.{js,ts,jsx,tsx,mdx}",
    "./src/components/**/*.{js,ts,jsx,tsx,mdx}",
    "./src/app/**/*.{js,ts,jsx,tsx,mdx}",
  ],
  theme: {
    
    extend: {
      colors: {
        background: "var(--background)",
        foreground: "var(--foreground)",
        mainColor: "#24211E",
        laranja:"#FF7F39"
      },
      boxShadow: {
        'custom': '10px 10px 0 rgba(255, 127, 57, 1)',
      },
      backgroundImage: {
        'home-1': 'url(/images/home-1.jpg)',
        'home-2': 'url(/images/home-2.jpg)',
        'home-3': 'url(/images/home-3.jpg)',
        'precos': 'url(/images/precos.jpg)',
        'cadastro': 'url(/images/cadastro.jpg)'
      },
      fontFamily: {
        'julius': ["'Julius Sans One'", "sans serif"],
        'archivo-black': ["'Archivo Black'", "sans-serif"],
        'archivo': ["'Archivo'", "sans-serif"]
      },
      animation: {
        fadeIn: "fadeIn 1s ease-in-out",
      },
      keyframes: {
        fadeIn: {
          "0%": { opacity: "0" },
          "100%": { opacity: "1" },
        },
      },
    },
  },
  plugins: [],
} satisfies Config;
