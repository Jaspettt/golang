const url = 'http://localhost:4000'

indexRender = async() => {
    try {
        let response = await fetch(url + '/api/v1/vinyls?limit=5&offset=0', {
            method: "GET",
            mode: "cors"
        })
        if (!response.ok) {
            throw new Error("error")
        }
        const data = await response.json()
        const offers = document.getElementById("offers")
        data.content.forEach(elem => {
            const div = document.createElement('div');
            div.classList.add('grid', 'items', 'bg-green-700', 'px-5', 'py-2', 'relative');
            div.innerHTML = `
                <h5 class="text-xl">${elem.title}</h5>
                <p class="text-sm">Artist: <span>${elem.artist}</span></p>
                <p class="text-sm">Release Date: <span>${elem.releasedate}</span>L</p>
                <p class="text-sm">Price: <span>${elem.price}</span></p>
                <p class="text-sm">Rating: <span>${elem.rating}</span></p>
                <button class="absolute bottom-0 right-0 text-white bg-blue-500 hover:bg-blue-700 px-7 py-2">Rent</button>
            `;
            offers.appendChild(div);
        });
    } catch(error) {
        console.error(error)
    }
}

window.onload = () => {
    indexRender()
}