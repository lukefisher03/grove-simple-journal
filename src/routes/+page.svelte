<script>
	import Icon from '@iconify/svelte';
	import Tiptap from '$lib/Tiptap.svelte';

	console.log('Hello');
</script>

<div class="flex h-screen p-0">
	<div class="flex flex-2  bg-[#3b3c41] text-[#88898d]">
		<div class="flex h-min flex-col px-3 mt-10">
			<div class="flex flex-col rounded-sm">
                <div class="flex rounded-sm p-1">
                    <Icon icon="ic:round-expand-more" width="24" height="24" />
                    <p>Tags</p>
                </div>
				<div class="flex flex-col ml-10 text-[#dfdfdf]">
					<div class="file-item flex p-1 rounded-sm font-bold">
						<Icon icon="" width="24" height="24" />
						<p>Testing</p>
					</div>
				</div>
			</div>
			<div class="file-item flex p-1">
				<Icon icon="" width="24" height="24" />
				<p>Cabinet 1</p>
			</div>
			<div class="file-item flex p-1">
				<Icon icon="" width="24" height="24" />
				<p>Cabinet 2</p>
			</div>
			<div class="file-item flex p-1">
				<Icon icon="" width="24" height="24" />
				<p>Cabinet 3</p>
			</div>
		</div>
	</div>
    <div class="flex flex-4 flex-col border-r-1 border-[#eaeaea] bg-white gap-6 overflow-y-scroll no-scrollbar">
        <h4 class="p-3 px-5 border-b-1 border-[#eaeaea] shadow-xs ml-2 sticky top-0 bg-[#fcfcfc]">Testing</h4>
        <div class="flex flex-col border-l-5 border-transparent p-2 mx-3">
            <h4>This is a neat article</h4>
            <p>Lorem ipsum dolor sit amet consectetur adipisicing elit. Quam nemo tempora distinctio, amet facere quos commodi unde, sit et itaque...</p>
        </div>
        <div class="flex flex-col mx-3 p-2 border-l-5 border-emerald-800 rounded-md bg-[#f5f5f5]">
            <h4>I love Svelte!</h4>
            <p>Lorem ipsum dolor sit amet consectetur adipisicing elit. Quam nemo tempora distinctio, amet facere quos commodi unde, sit et itaque...</p>
        </div>
        <div class="flex flex-col border-l-5 border-transparent p-2 mx-3">
            <h4>Cool design ideas</h4>
            <p>Lorem ipsum dolor sit amet consectetur adipisicing elit. Quam nemo tempora distinctio, amet facere quos commodi unde, sit et itaque...</p>
        </div>
        <div class="flex flex-col border-l-5 border-transparent mx-3 p-2">
            <h4>Don't pay for notes apps!</h4>
            <p>Lorem ipsum dolor sit amet consectetur adipisicing elit. Quam nemo tempora distinctio, amet facere quos commodi unde, sit et itaque...</p>
        </div>
        <div class="flex flex-col border-l-5 border-transparent mx-3 p-2">
            <h4>This is the article list!</h4>
            <p>Lorem ipsum dolor sit amet consectetur adipisicing elit. Quam nemo tempora distinctio, amet facere quos commodi unde, sit et itaque...</p>
        </div>
        <div class="flex flex-col border-l-5 border-transparent mx-3 p-2">
            <h4>This is the article list!</h4>
            <p>Lorem ipsum dolor sit amet consectetur adipisicing elit. Quam nemo tempora distinctio, amet facere quos commodi unde, sit et itaque...</p>
        </div>
        <div class="flex flex-col border-l-5 border-transparent mx-3 p-2">
            <h4>This is the article list!</h4>
            <p>Lorem ipsum dolor sit amet consectetur adipisicing elit. Quam nemo tempora distinctio, amet facere quos commodi unde, sit et itaque...</p>
        </div>
        <div class="flex flex-col border-l-5 border-transparent mx-3 p-2">
            <h4>This is the article list!</h4>
            <p>Lorem ipsum dolor sit amet consectetur adipisicing elit. Quam nemo tempora distinctio, amet facere quos commodi unde, sit et itaque...</p>
        </div>
        <div class="flex flex-col border-l-5 border-transparent mx-3 p-2">
            <h4>This is the article list!</h4>
            <p>Lorem ipsum dolor sit amet consectetur adipisicing elit. Quam nemo tempora distinctio, amet facere quos commodi unde, sit et itaque...</p>
        </div>
        <div class="flex flex-col border-l-5 border-transparent mx-3 p-2">
            <h4>This is the article list!</h4>
            <p>Lorem ipsum dolor sit amet consectetur adipisicing elit. Quam nemo tempora distinctio, amet facere quos commodi unde, sit et itaque...</p>
        </div>
        <div class="flex flex-col border-l-5 border-transparent mx-3 p-2">
            <h4>This is the article list!</h4>
            <p>Lorem ipsum dolor sit amet consectetur adipisicing elit. Quam nemo tempora distinctio, amet facere quos commodi unde, sit et itaque...</p>
        </div>
    </div>
	<div class="tiptap-container h-full flex-13">
		<Tiptap />
	</div>
</div>
