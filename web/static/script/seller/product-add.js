const name = document.getElementById('name');
const price = document.getElementById('price');
const description = document.getElementById('description');
const image = document.getElementById('image');
const category = document.getElementById('category');
const stock = document.getElementById('stock');
const weight = document.getElementById("weight")

const file = document.getElementById('file');

const form = document.getElementById('form');

file.addEventListener('change', async (e) => { 
	e.preventDefault();
	const formData = new FormData();
	formData.append('image', file.files[0]);

	const res = await fetch('/api/product/image', {
		method: 'POST',
		body: formData,
	})

	if (!res.ok) {
		console.log(res.statusText);
		return;
	}

	const data = await res.json();
	console.log(data);

	image.value = data.value.image;
});

form.addEventListener('submit', async (e) => {
	e.preventDefault()
	const res = await fetch('/api/product', {
		method: 'POST',
		headers: {
			'Content-Type': 'application/json'
		},
		body: JSON.stringify({
			name: name.value,
			price: Number(price.value), 
			description: description.value,
			image: image.value,
			category: category.value,
			stock: Number(stock.value),
			weight: Number(weight.value)
		})
	})

	if (!res.ok) {
		console.log(res.statusText);
		return
	}

	const data = await res.json()

	console.log(data);
});