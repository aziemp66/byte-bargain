const list = document.createElement("ul");

for (let i = 0; i < 5; i++) {
	const item = document.createElement("li");
	item.append(`Item ${i}`);
	list.append(item);
}

document.querySelector("body").append(list);