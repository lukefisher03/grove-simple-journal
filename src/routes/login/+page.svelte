<script>
	console.log('Hello');
</script>

<div class="flex h-screen items-center justify-center">
    <div class="flex flex-col items-center min-w-md rounded-xl border-1 border-[#e4e4e4] py-15 bg-white shadow-sm">
        <div class="flex flex-col gap-8 mx-w-sm">
            <div class="justify-self-start self-start">
                <h1 class="text-2xl">Sign into <span class="font-bold italic">Noted</span></h1>
            </div>
            <div class="flex flex-col">
                <div>
                    <label for="username">Email</label>
                </div>
                <div>
                    <input type="text" name="username" placeholder="hello@email.com" class="bg-[#f3f4f6] inset-shadow-lg rounded-md min-w-sm p-3">
                </div>
            </div>
            <div class="flex flex-col gap-1">
                <div>
                    <label for="username">Password</label>
                </div>
                <div>
                    <input type="password" name="username" placeholder="Password" class="bg-[#f3f4f6] inset-shadow-md rounded-md min-w-sm p-2">
                </div>
                <a href="/" class="italic underline text-xs text-[#9e9fa2]">Forgot password?</a>
            </div>
            <div class="flex flex-col gap-4">
                <input type="submit" value="Continue" class="p-2 px-8 bg-emerald-800 rounded-md shadow-2xs text-white">
                <div class="or-divider">or</div>
                <input type="submit" value="Sign up" class="p-2 px-8 border-1 border-[#e4e4e4] rounded-md shadow-2xs">
            </div>
        </div>
	</div>
</div>
