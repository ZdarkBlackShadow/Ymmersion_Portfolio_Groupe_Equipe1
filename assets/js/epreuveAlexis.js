function openTab(tabId) {
    // Masquer tous les contenus d'onglet
    const tabContents = document.querySelectorAll('.tab-content');
    tabContents.forEach(tab => tab.classList.remove('active-content'));

    // Retirer la classe active de tous les onglets
    const tabs = document.querySelectorAll('.tab');
    tabs.forEach(tab => tab.classList.remove('active-tab'));

    // Afficher le contenu de l'onglet sélectionné et activer l'onglet
    document.getElementById(tabId).classList.add('active-content');
    const activeTab = Array.from(tabs).find(tab => tab.classList.contains(tabId));
    activeTab.classList.add('active-tab');

    // Changer la couleur de fond de chaque onglet
    tabs.forEach(tab => {
        tab.style.backgroundColor = "#fd9595"; // Réinitialise la couleur de fond pour tous
    });

    // Appliquer une nouvelle couleur de fond pour l'onglet actif
    activeTab.style.backgroundColor = "#9cfea2"; // Changer la couleur de fond de l'onglet actif
}
