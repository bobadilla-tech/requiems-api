import { Controller } from "@hotwired/stimulus";

// Global search controller for navbar
// Searches across APIs, Examples, and Pages
export default class extends Controller {
  static targets = ["input", "dropdown", "results", "empty"];
  static values = {
    searchData: Object,
    minLength: { type: Number, default: 2 },
  };

  connect() {
    // Close dropdown when clicking outside
    this.handleClickOutside = this.handleClickOutside.bind(this);
    document.addEventListener("click", this.handleClickOutside);

    // Initialize keyboard navigation
    this.selectedIndex = -1;
  }

  disconnect() {
    document.removeEventListener("click", this.handleClickOutside);
  }

  // Search and filter content
  search(event) {
    const query = this.inputTarget.value.toLowerCase().trim();

    // Hide dropdown if query is too short
    if (query.length < this.minLengthValue) {
      this.hideDropdown();
      return;
    }

    // Filter all content types
    const apiResults = this.filterContent(this.searchDataValue.apis, query);
    const exampleResults = this.filterContent(
      this.searchDataValue.examples,
      query,
    );
    const pageResults = this.filterContent(this.searchDataValue.pages, query);

    const allResults = [
      ...apiResults,
      ...exampleResults,
      ...pageResults,
    ];

    // Display results
    this.displayResults(allResults, query);
    this.showDropdown();
  }

  // Filter content by query
  filterContent(items, query) {
    return items.filter((item) => {
      const titleMatch = item.title.toLowerCase().includes(query);
      const descMatch = item.description.toLowerCase().includes(query);
      const tagsMatch = item.tags &&
        item.tags.some((tag) => tag.toLowerCase().includes(query));
      const categoryMatch = item.category &&
        item.category.toLowerCase().includes(query);
      const techMatch = item.technologies &&
        item.technologies.some((tech) => tech.toLowerCase().includes(query));

      return titleMatch || descMatch || tagsMatch || categoryMatch || techMatch;
    });
  }

  // Display filtered results
  displayResults(results, query) {
    const maxResults = 8;
    const displayResults = results.slice(0, maxResults);

    if (displayResults.length === 0) {
      this.showEmpty();
      return;
    }

    this.hideEmpty();

    // Build results HTML grouped by type
    const html = displayResults.map((item, index) =>
      this.buildResultHTML(item, index, query)
    ).join("");

    // Add "View all API results" footer if there are API matches
    const apiCount = results.filter((r) => r.type === "api").length;
    const footer = apiCount > 0
      ? `
      <div class="border-t border-gray-200 dark:border-gray-600 px-4 py-3 bg-gray-50 dark:bg-gray-900/80">
        <a href="/apis?q=${encodeURIComponent(query)}"
           class="text-sm text-blue-600 hover:text-blue-700 dark:text-blue-400 dark:hover:text-blue-300 font-medium">
          View all ${apiCount} API result${apiCount !== 1 ? "s" : ""} →
        </a>
      </div>
    `
      : "";

    this.resultsTarget.innerHTML = html + footer;
  }

  // Build HTML for a single result
  buildResultHTML(item, index, query) {
    const badge = this.getTypeBadge(item.type);
    const icon = this.getIcon(item);
    const subtitle = this.getSubtitle(item);

    return `
      <a href="${item.url}"
         class="navbar-search-result flex items-start px-4 py-3 hover:bg-gray-50 dark:hover:bg-gray-700/80 transition-colors border-l-4 border-transparent hover:border-blue-500 dark:hover:border-blue-400"
         data-index="${index}">
        <div class="flex-shrink-0 mr-3 text-2xl">
          ${icon}
        </div>
        <div class="flex-1 min-w-0">
          <div class="flex items-center justify-between mb-1">
            <h4 class="text-sm font-semibold text-gray-900 dark:text-gray-100 truncate">
              ${this.highlightQuery(item.title, query)}
            </h4>
            ${badge}
          </div>
          <p class="text-xs text-gray-600 dark:text-gray-300 line-clamp-2 mb-1">
            ${this.highlightQuery(item.description, query)}
          </p>
          ${subtitle}
        </div>
      </a>
    `;
  }

  // Get type badge HTML
  getTypeBadge(type) {
    const badges = {
      api:
        '<span class="ml-2 inline-flex shrink-0 items-center px-2 py-0.5 rounded text-xs font-medium bg-blue-100 text-blue-800 dark:bg-blue-950/80 dark:text-blue-200">API</span>',
      example:
        '<span class="ml-2 inline-flex shrink-0 items-center px-2 py-0.5 rounded text-xs font-medium bg-purple-100 text-purple-800 dark:bg-purple-950/80 dark:text-purple-200">Example</span>',
      page:
        '<span class="ml-2 inline-flex shrink-0 items-center px-2 py-0.5 rounded text-xs font-medium bg-gray-100 text-gray-800 dark:bg-gray-600 dark:text-gray-100">Page</span>',
    };
    return badges[type] || "";
  }

  // Get icon for result
  getIcon(item) {
    if (item.type === "api") {
      return item.category_icon || "🔌";
    } else if (item.type === "example") {
      return "💡";
    } else if (item.type === "page") {
      return item.icon || "📄";
    }
    return "📄";
  }

  // Get subtitle for result
  getSubtitle(item) {
    if (item.type === "api") {
      return `
        <div class="flex items-center text-xs text-gray-500 dark:text-gray-400 space-x-2">
          <span>${item.category}</span>
          <span>•</span>
          <span>${item.endpoints_count} ${
        item.endpoints_count === 1 ? "endpoint" : "endpoints"
      }</span>
        </div>
      `;
    } else if (item.type === "example") {
      const techBadges = item.technologies.slice(0, 3).map((tech) =>
        `<span class="inline-flex items-center px-1.5 py-0.5 rounded text-xs bg-gray-100 text-gray-700 dark:bg-gray-600 dark:text-gray-200">${tech}</span>`
      ).join(" ");
      return `
        <div class="flex items-center space-x-1 mt-1">
          ${techBadges}
        </div>
      `;
    } else if (item.type === "page") {
      return `<span class="text-xs text-gray-500 dark:text-gray-400">${
        item.category || "Resource"
      }</span>`;
    }
    return "";
  }

  // Highlight matching query text
  highlightQuery(text, query) {
    if (!text) return "";
    const regex = new RegExp(`(${this.escapeRegex(query)})`, "gi");
    return text.replace(
      regex,
      '<mark class="bg-yellow-200 text-gray-900 dark:bg-yellow-500/40 dark:text-gray-100 font-medium rounded px-0.5">$1</mark>',
    );
  }

  // Escape special regex characters
  escapeRegex(string) {
    return string.replace(/[.*+?^${}()|[\]\\]/g, "\\$&");
  }

  // Show empty state
  showEmpty() {
    this.resultsTarget.innerHTML = "";
    this.emptyTarget.classList.remove("hidden");
  }

  // Hide empty state
  hideEmpty() {
    this.emptyTarget.classList.add("hidden");
  }

  // Show dropdown
  showDropdown() {
    this.dropdownTarget.classList.remove("hidden");
  }

  // Hide dropdown
  hideDropdown() {
    this.dropdownTarget.classList.add("hidden");
    this.selectedIndex = -1;
  }

  // Handle click outside to close dropdown
  handleClickOutside(event) {
    if (!this.element.contains(event.target)) {
      this.hideDropdown();
    }
  }

  // Clear search
  clear(event) {
    event.preventDefault();
    event.stopPropagation();
    this.inputTarget.value = "";
    this.hideDropdown();
    this.inputTarget.focus();
  }

  // Keyboard navigation
  handleKeydown(event) {
    const results = this.resultsTarget.querySelectorAll(
      ".navbar-search-result",
    );

    switch (event.key) {
      case "Escape":
        event.preventDefault();
        this.hideDropdown();
        this.inputTarget.blur();
        break;

      case "ArrowDown":
        event.preventDefault();
        this.selectedIndex = Math.min(
          this.selectedIndex + 1,
          results.length - 1,
        );
        this.updateSelection(results);
        break;

      case "ArrowUp":
        event.preventDefault();
        this.selectedIndex = Math.max(this.selectedIndex - 1, -1);
        this.updateSelection(results);
        break;

      case "Enter":
        event.preventDefault();
        if (this.selectedIndex >= 0 && results[this.selectedIndex]) {
          results[this.selectedIndex].click();
        }
        break;
    }
  }

  // Update visual selection
  updateSelection(results) {
    results.forEach((result, index) => {
      if (index === this.selectedIndex) {
        result.classList.add(
          "bg-gray-100",
          "dark:bg-gray-700",
          "border-blue-500",
          "dark:border-blue-400",
        );
        result.classList.remove("border-transparent");
        result.scrollIntoView({ block: "nearest", behavior: "smooth" });
      } else {
        result.classList.remove(
          "bg-gray-100",
          "dark:bg-gray-700",
          "border-blue-500",
          "dark:border-blue-400",
        );
        result.classList.add("border-transparent");
      }
    });
  }

  // Focus search (called from keyboard shortcut controller)
  focus() {
    this.inputTarget.focus();
    this.inputTarget.select();
  }
}
