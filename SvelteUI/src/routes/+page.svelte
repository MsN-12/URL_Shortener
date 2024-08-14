<section class="flex-1  w-full  px-3 antialiased bg-gradient-to-br from-purple-900 via-black to-purple-900 lg:px-6 min-h-screen">
    <div class="mx-auto max-w-7xl">
        <div class="container px-6 mx-auto md:text-center md:px-1 mt-10">
            <h1 class="text-4xl font-extrabold leading-10 tracking-tight text-white sm:text-5xl md:text-6xl xl:text-6xl"><span class="block">Welcome to Kootah Link</span> </h1>
            <h1 class="text-xl font-bold leading-10 tracking-tight text-white sm:text-5xl md:text-6xl xl:text-3xl mt-5"><span class="block">Your Trusted URL Shortener</span> </h1>
            <div class="flex flex-col relative items-end mx-auto mt-12 overflow-hidden text-left  rounded-md md:max-w-4xl md:text-center">
                <textarea bind:value={longUrl} name="LongUrl" placeholder="Enter your long URL ..." class=" w-full h-36 px-6 py-2 font-medium text-6 text-gray-700 focus:outline-none"></textarea>
                <button on:click={createShortUrl} type="button" class="mt-4 px-4 py-3 rounded-lg text-base font-bold leading-6 text-white transition duration-150 ease-in-out bg-purple-600 border border-transparent hover:bg-purple-500 focus:outline-none active:bg-purple-600" data-primary="gray-500">
                    Submit
                </button>
            </div>
            {#if shortUrl}
                <div class="container mx-auto px-6 py-4 mt-6 text-center">
                    <h2 class="text-2xl font-semibold text-white mb-4">Here's your shortened URL:</h2>
                    <div class="inline-block text-lg font-medium py-2 px-4 text-white bg-purple-900 rounded-lg shadow-md hover:bg-purple-800 transition ease-in-out duration-150">
                        <a href={shortUrl} class="hover:text-gray-300" target="_blank">{shortUrl}</a>
                    </div>
                    <button on:click={copyToClipboard} class="ml-2 bg-blend-color-dodge inline-block text-lg font-medium py-2 px-4 text-white bg-purple-900 rounded-lg shadow-md hover:bg-purple-800 transition ease-in-out duration-150">
                        Copy <i class="fa fa-copy"></i>
                    </button>
                </div>
            {/if}
            {#if error}
                <div class="text-red-500">{error}</div>
            {/if}
        </div>
    </div>
</section>

<script>
    function copyToClipboard() {
        navigator.clipboard.writeText(shortUrl);
    }
    let longUrl = '';
    let shortUrl = '';
    let error = '';

    async function createShortUrl() {
        try {
            const response = await fetch('http://localhost:3000/api/create-url', {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json',
                },
                body: JSON.stringify({ long_url: longUrl }),
            });

            if (!response.ok) {
                throw new Error(`HTTP error! status: ${response.status}`);
            }

            const data = await response.json();
            shortUrl = "http://localhost:3000/"+data.short_url;
        } catch (e) {
            error = (e instanceof Error) ? e.message : String(e);
        }
    }
</script>