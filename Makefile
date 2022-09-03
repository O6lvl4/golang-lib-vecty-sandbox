setup:
	echo setup
gen-css:
	npx tailwindcss-cli@latest build ./res/css/import.css -o ./res/css/tailwind.css
