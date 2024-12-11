function openTab(evt, tabName) {
    var i, tabContent, tabLinks;
    tabContent = document.getElementsByClassName("tab-content");
    for (i = 0; i < tabContent.length; i++) {
        tabContent[i].style.display = "none";
    }
    tabLinks = document.getElementsByClassName("tab");
    for (i = 0; i < tabLinks.length; i++) {
        tabLinks[i].className = tabLinks[i].className.replace(" active", "");
    }
    document.getElementById(tabName).style.display = "block";
    evt.currentTarget.className += " active";
}

const cursor = document.createElement('div');
cursor.classList.add('custom-cursor');
document.body.appendChild(cursor);

document.addEventListener('mousemove', (e) => {
    cursor.style.left = e.clientX + 'px';
    cursor.style.top = e.clientY + 'px';

    // Check if the element under cursor is clickable
    const elementUnderCursor = document.elementFromPoint(e.clientX, e.clientY);
    if (elementUnderCursor) {
        const isClickable = (
            elementUnderCursor.matches('a, button, [role="button"], input, .clickable') ||
            elementUnderCursor.closest('a, button, [role="button"], input, .clickable')
        );

        // Set opacity based on whether element is clickable
        cursor.style.opacity = isClickable ? '0' : '1';
    }
});

document.addEventListener('DOMContentLoaded', function () {
    const tooltip = document.getElementById('tooltip');
    const buttons = document.querySelectorAll('.search-button, .details-button, .home-link');
    buttons.forEach(button => {
        button.addEventListener('mouseenter', (event) => {
            tooltip.textContent = button.getAttribute('data-tooltip');
            tooltip.style.display = 'block';
            // Get the position of the button and set the tooltip below it
            const rect = event.target.getBoundingClientRect();
            tooltip.style.left = `${rect.left + window.scrollX + rect.width / 2 - tooltip.offsetWidth / 2}px`;
            tooltip.style.top = `${rect.bottom + window.scrollY + 5}px`;
        });

        button.addEventListener('mouseleave', () => {
            tooltip.style.display = 'none';
        });
    });
});

document.addEventListener('keydown', function(event) {
    // Check if the CTRL key and the 'B' key are pressed
    if (event.ctrlKey && event.key.toLowerCase() === 'b') {
        event.preventDefault(); // Prevent any default action for CTRL + B
        window.history.back(); // Go to the previous page
    }
});
document.addEventListener('DOMContentLoaded', () => {
    const searchInput = document.getElementById('search-input');
    const suggestionsContainer = document.getElementById('search-suggestions');

    searchInput.addEventListener('input', async (e) => {
        const query = e.target.value.trim();
        
        try {
            const response = await fetch(`/search/suggestions?q=${encodeURIComponent(query)}`);
            const suggestions = await response.json();

            // Clear previous suggestions
            suggestionsContainer.innerHTML = '';

            // Render new suggestions
            suggestions.forEach(suggestion => {
                const suggestionElement = document.createElement('div');
                suggestionElement.classList.add('suggestion-item');
                suggestionElement.innerHTML = `
                    <span class="suggestion-text">${suggestion.Value}</span>
                    <span class="suggestion-type">${suggestion.Type}</span>
                `;
                
                // Add click event to fill search input
                suggestionElement.addEventListener('click', () => {
                    searchInput.value = suggestion.Value;
                    suggestionsContainer.innerHTML = '';
                    // Optional: trigger search
                });

                suggestionsContainer.appendChild(suggestionElement);
            });
        } catch (error) {
            console.error('Error fetching suggestions:', error);
        }
    });

    document.addEventListener('click', (e) => {
        if (!searchInput.contains(e.target) && !suggestionsContainer.contains(e.target)) {
            suggestionsContainer.innerHTML = '';
        }
    });
});