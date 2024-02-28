export function initBillMaintenance(page) {
    $.ajax({
        url: 'http://127.0.0.1:8228/list',
        type: 'POST',
        data: JSON.stringify({
            page: page,
            page_size: 10
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
