export default defineNuxtConfig({
	app: {
		head: {
			charset: 'UTF-8',
			viewport: 'width=device-width, initial-scale=1.0, minimum-scale=1.0, maximum-scale=1.0, user-scalable=no',
			link: [
				{
					rel: 'icon',
					href: '/favicon.ico',
					sizes: '32x32',
				},
				{
					rel: 'icon',
					href: '/favicon.svg',
					type: 'image/svg+xml',
				},
				{
					rel: 'apple-touch-icon',
					href: '/apple-touch-icon.png',
				},
				{
					rel: 'manifest',
					href: '/manifest.json',
				},
			],
		},
	},
	modules: [
		'@nuxt/ui',
	],
	css: [
		'~/assets/css/app.css',
	],
});
