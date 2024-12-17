document.addEventListener('DOMContentLoaded', () => {
    // Cursor customization
    const cursor = document.createElement('div');
    cursor.classList.add('custom-cursor');
    document.body.appendChild(cursor);

    function isElementClickable(element) {
        return element.matches('a, button, [role="button"], input, .clickable') ||
               element.closest('a, button, [role="button"], input, .clickable');
    }

    document.addEventListener('mousemove', (e) => {
        cursor.style.left = `${e.clientX}px`;
        cursor.style.top = `${e.clientY}px`;
        const elementUnderCursor = document.elementFromPoint(e.clientX, e.clientY);
        cursor.style.opacity = elementUnderCursor && isElementClickable(elementUnderCursor) ? '0' : '1';
    });

    // Tooltip functionality
    const tooltip = document.getElementById('tooltip');
    document.querySelectorAll('.search-button, .details-button, .home-link').forEach(button => {
        button.addEventListener('mouseenter', (event) => {
            tooltip.textContent = event.target.getAttribute('data-tooltip');
            tooltip.style.display = 'block';
            const rect = event.target.getBoundingClientRect();
            tooltip.style.left = `${rect.left + window.scrollX + rect.width / 2 - tooltip.offsetWidth / 2}px`;
            tooltip.style.top = `${rect.bottom + window.scrollY + 5}px`;
        });

        button.addEventListener('mouseleave', () => {
            tooltip.style.display = 'none';
        });
    });

    // Search suggestions
    const searchInput = document.getElementById('search-input');
    const suggestionsContainer = document.getElementById('search-suggestions');

    searchInput.addEventListener('input', async (e) => {
        const query = e.target.value.trim();
        if (!query) {
            suggestionsContainer.innerHTML = '';
            return;
        }

        try {
            const response = await fetch(`/search/suggestions?q=${encodeURIComponent(query)}`);
            const suggestions = await response.json();
            suggestionsContainer.innerHTML = '';
            suggestions.forEach(suggestion => {
                const suggestionElement = document.createElement('div');
                suggestionElement.classList.add('suggestion-item');
                suggestionElement.innerHTML = `
                   <a href="/artists/?id=${suggestion.ArtistID}"> <span class="suggestion-text">${suggestion.Value} - ${suggestion.Type} of ${suggestion.Name}</span>
                `;
                suggestionElement.addEventListener('click', () => {
                    searchInput.value = suggestion.Value;
                    suggestionsContainer.innerHTML = '';
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

    // CTRL + B to go back
    document.addEventListener('keydown', (event) => {
        if (event.ctrlKey && event.key.toLowerCase() === 'b') {
            event.preventDefault();
            window.history.back();
        }
    });
});
