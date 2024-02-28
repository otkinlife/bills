function initBillMaintenance(page, filters = {}) {
    $.ajax({
        url: 'http://127.0.0.1:8228/list',
        type: 'POST',
        data: JSON.stringify({
            page: page,
            page_size: 10,
            status: filters.status,
            channel: filters.channel,
            tag: filters.tag,
            transaction_type: filters.transaction_type,
            charge_name: filters.name,
            date_range:[filters.date_range[0],filters.date_range[1]],
            value_range:[filters.value_range[0],filters.value_range[1]]
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
            var startPage = Math.max(1, page - 2);
            var endPage = Math.min(data.total, startPage + 4);
            startPage = Math.max(1, endPage - 4);

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
                    initBillMaintenance(page - 1);
                });
            } else {
                previousPage.addClass('disabled');
                previousPage.find('a').off('click');
            }

            if (page < data.total) {
                nextPage.removeClass('disabled');
                nextPage.find('a').off('click').on('click', function (e) {
                    e.preventDefault();
                    initBillMaintenance(page + 1);
                });
            } else {
                nextPage.addClass('disabled');
                nextPage.find('a').off('click');
            }

            // Handle the go-button click event
            $('#go-button').off('click').on('click', function () {
                var gotoPage = parseInt($('#page-input').val());
                if (gotoPage >= 1 && gotoPage <= data.total) {
                    initBillMaintenance(gotoPage);
                } else {
                    alert('请输入有效的页码');
                }
            });
        },
        error: function (xhr, status, error) {
            alert('获取账单数据失败：' + error);
        }
    });

    function createPageLinkHandler(page) {
        return function (e) {
            e.preventDefault();
            initBillMaintenance(page);
        };
    }
}

export function initFilterForm(page) {
    // Fetch the filter options from the server
    fetchFilterOptions().then(fillFilterForm);

    // Handle the form submission
    const filterForm = document.getElementById('filter-form');
    filterForm.addEventListener('submit', function (e) {
        e.preventDefault();

        // Get the filter values
        const filters = {
            date_range: [document.getElementById('start_date').value, document.getElementById('end_date').value],
            value_range: [document.getElementById('min_value').value, document.getElementById('max_value').value],
            channel: document.getElementById('channel').value,
            tag: document.getElementById('tag').value,
            transaction_type: document.getElementById('transaction_type').value,
            name: document.getElementById('charge_name').value,
            status: document.getElementById('status').value
        };

        // Call the function to filter the table data
        filterTableData(page, filters);
    });
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
        console.log(key);
        let select = document.getElementById(key);
        data.data[key].forEach(function (option) {
            let optionElement = document.createElement('option');
            optionElement.textContent = option;
            select.appendChild(optionElement);
        });
    }
}

function filterTableData(page, filters) {
    initBillMaintenance(page, filters);
}
