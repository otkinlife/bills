function formatDate(date) {
    let day = ('0' + date.getDate()).slice(-2);
    let month = ('0' + (date.getMonth() + 1)).slice(-2);
    let year = date.getFullYear();

    return year + '-' + month + '-' + day;
}

function createOptionElement(text) {
    let optionElement = document.createElement('option');
    optionElement.textContent = text;
    return optionElement;
}

function getFilterValues() {
    return  {
        date_range: [document.getElementById('start_date').value, document.getElementById('end_date').value],
        value_range: [parseFloat(document.getElementById('min_value').value) || 0, parseFloat(document.getElementById('max_value').value) || 0],
        channel: document.getElementById('channel').value || "不限",
        tag: document.getElementById('tag').value || "不限",
        transaction_type: document.getElementById('transaction_type').value || "不限",
        name: document.getElementById('charge_name').value || "不限",
        status: document.getElementById('status').value || "不限"
    };
}

function fetchFilterOptions() {
    return new Promise(function (resolve, reject) {
        $.ajax({
            url: '/list_dict',
            method: 'GET',
            success: resolve,
            error: reject
        });
    });
}

function fillFilterForm(data) {
    // Fill the select elements with the received options
    for (let key in data.data) {
        let select = document.getElementById(key);
        select.innerHTML = '';  // Clear the old options
        select.appendChild(createOptionElement('不限')); // Add the default option
        data.data[key].forEach(function (option) {
            select.appendChild(createOptionElement(option));
        });
    }
}

function handleFilterForm(page) {
    // Fetch the filter options from the server
    fetchFilterOptions().then(fillFilterForm);

    // Set the default date range to the last month
    let endDate = new Date();
    let startDate = new Date();
    startDate.setMonth(startDate.getMonth() - 1);
    document.getElementById('start_date').value = formatDate(startDate);
    document.getElementById('end_date').value = formatDate(endDate);

    // Get the filter values and initialize the table
    let filters = getFilterValues();
    initBillMaintenance(page, filters);

    // Handle the form submission
    const filterForm = document.getElementById('filter-form');
    filterForm.addEventListener('submit', function (e) {
        e.preventDefault();
        filters = getFilterValues();
        initBillMaintenance(page, filters);
    });
}

function initBillMaintenance(page, filters = {}) {
    $.ajax({
        url: 'http://127.0.0.1:8228/list',
        type: 'POST',
        data: JSON.stringify({
            page: page,
            page_size: 10,
            status: filters.status || null,
            channel: filters.channel || null,
            tag: filters.tag || null,
            transaction_type: filters.transaction_type || null,
            charge_name: filters.name || null,
            date_range: filters.date_range || [null, null],
            value_range: filters.value_range || [null, null]
        }),
        contentType: 'application/json; charset=utf-8',
        success: function (data) {
            var tbody = $('#bill-table tbody');
            tbody.empty();

            $.each(data.list, function (i, bill) {
                var tr = $('<tr>');
                tr.append($('<td class="text-overflow">').text(bill.Goods));
                tr.append($('<td class="text-overflow">').text(bill.Value));
                tr.append($('<td class="text-overflow">').text(bill.Tag));
                tr.append($('<td class="text-overflow">').text(bill.TransactionType));
                tr.append($('<td class="text-overflow">').text(bill.ChargeTime));
                tr.append($('<td class="text-overflow">').text(bill.Status));
                tbody.append(tr);
            });

            // Update the pagination
            var pagination = $('.pagination');
            pagination.children().not('#previous-page, #next-page, #go-page').remove(); // Remove old page links
            var previousPage = $('#previous-page');
            var nextPage = $('#next-page');
            var pageCount = Math.floor(data.total / 10);
            if (data.total % 10 !== 0) {
                pageCount += 1;
            }
            var startPage = Math.max(1, page - 2);
            var endPage = Math.min(pageCount, startPage + 4);

            for (var i = startPage; i <= endPage; i++) {
                var li = $('<li>').addClass('page-item');
                if (i === page) {
                    li.addClass('active');
                }
                var a = $('<a>').addClass('page-link').text(i).attr('href', '#');
                a.on('click', createPageLinkHandler(i));
                li.append(a);
                nextPage.before(li); // Insert li before #next-page
            }

            if (page > 1) {
                previousPage.removeClass('disabled');
                previousPage.find('a').off('click').on('click', function (e) {
                    e.preventDefault();
                    handleFilterForm(page - 1);
                });
            } else {
                previousPage.addClass('disabled');
                previousPage.find('a').off('click');
            }

            if (page < data.total) {
                nextPage.removeClass('disabled');
                nextPage.find('a').off('click').on('click', function (e) {
                    e.preventDefault();
                    handleFilterForm(page + 1);
                });
            } else {
                nextPage.addClass('disabled');
                nextPage.find('a').off('click');
            }
        },
        error: function (xhr, status, error) {
            alert('获取账单数据失败：' + error);
        }
    });

    function createPageLinkHandler(page) {
        return function (e) {
            e.preventDefault();
            handleFilterForm(page);
        };
    }
}


$(document).ready(function() {
    handleFilterForm(1);  // Initialize the first page
});