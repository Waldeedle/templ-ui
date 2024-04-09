/** @type {import('tailwindcss').Config} */
module.exports = {
    content: ['./ui/*{_templ.go,.templ}'],
    theme: {
        extend: {},
    },
    plugins: [
        require('@tailwindcss/forms'),
        require('@tailwindcss/typography'),
    ]
}