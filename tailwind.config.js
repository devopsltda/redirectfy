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
    extend: {
      colors: {
        "rblack": "#000000",
        "rdarker": "#773d04",
        "rdark": "#a65505",
        "rprimary": "#ed7a07",
        "rlight": "#f2a251",
        "rlighter": "#f8ca9c",
        "rlightest": "#fdf2e6",
      },
    },
  },
  plugins: [],
};

