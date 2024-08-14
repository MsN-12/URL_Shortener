<script>
    import "../app.css";
</script>
<head>
   <title >Kootah_Link.</title>
   <link rel="icon" type="image/x-icon" href="/static/favicon.ico"/>
   <meta charset="UTF-8"/>
   <meta name="viewport" content="width=device-width, initial-scale=1.0"/>
   <link rel="stylesheet" href="https://cdnjs.cloudflare.com/ajax/libs/font-awesome/4.7.0/css/font-awesome.min.css">
</head>
<body class="antialiased flex-col flex h-24">
   <section class="w-full px-8 text-purple-700 bg-white">
      <div class="container flex flex-col flex-wrap items-center justify-between py-6 mx-auto md:flex-row max-w-7xl">
         <div class="relative flex flex-col md:flex-row">
            <a href="/" class="flex items-center mb-5 font-medium text-purple-900 lg:w-auto lg:items-center lg:justify-center md:mb-0">
               <span class="mx-auto text-xl font-black leading-none text-purple-500 select-none hover:text-purple-400">Kootah Link<span class="text-indigo-600" data-primary="indigo-600">.</span></span>
            </a>
            <nav class="flex flex-wrap items-center mb-5 text-base md:mb-0 md:pl-8 md:ml-8 md:border-l md:border-gray-200">
               <a href="/" class="mr-5 font-medium leading-6 text-gray-600 hover:text-gray-900">Home</a>
               <a href="#_" class="mr-5 font-medium leading-6 text-gray-600 hover:text-gray-900">Features</a>
               <a href="#_" class="mr-5 font-medium leading-6 text-gray-600 hover:text-gray-900">Pricing</a>
               <a href="#_" class="mr-5 font-medium leading-6 text-gray-600 hover:text-gray-900">Blog</a>
            </nav>
         </div>
         <div class="inline-flex items-center space-x-6 lg:justify-end">
            <a href="/auth/register" class="text-base font-medium leading-6 text-purple-600 whitespace-no-wrap transition duration-150 ease-in-out hover:text-purple-900">
               Sign in
            </a>
            <a href="/auth/login" class="inline-flex items-center w-full px-6  py-3 text-sm font-medium leading-6 text-white bg-purple-600 md:w-auto md:rounded-2xl hover:bg-purple-500 focus:outline-none md:focus:ring-2 focus:ring-0 focus:ring-offset-2 focus:ring-purple-700">
               Sign up
            </a>
         </div>
      </div>
   </section>
</body>
<slot></slot>
