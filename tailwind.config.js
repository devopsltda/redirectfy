const defaultTheme = require('tailwindcss/defaultTheme');

/** @type {import('tailwindcss').Config} */
module.exports = {
  content: [
    "./cmd/web/**/*.templ",
  ],
  theme: {
    fontFamily: {
      'sans': ['"DM Sans"', ...defaultTheme.fontFamily.sans],
    },
    extend: {},
  },
  plugins: [],
};

