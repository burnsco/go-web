/** @type {import('tailwindcss').Config} */
module.exports = {
  content: [
    "./ui/html/*.tmpl"
  ],
  theme: {
    extend: {},
  },
  plugins: [
    require('@tailwindcss/forms'),
    require('@tailwindcss/typography'),
  ],
}

