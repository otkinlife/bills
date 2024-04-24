$(document).ready(function () {
    // Add an event listener to the navigation items
    $('.nav-item').on('click', function () {
        // Show the loading animation
        $('#loading').show();

        // Remove the 'active' class from the currently active item
        $('.nav-item.active').removeClass('active');

        // Add the 'active' class to the clicked item
        $(this).addClass('active');

        if (this.children[0].id === 'dashboard-link') {
            // If the dashboard link was clicked, show the dashboard iframe
            var dashboardUrl = 'http://dev.com:3000/d-solo/ddjojuc52rsowf/e694af-e587ba-e79c8b-e69dbf?orgId=1&from=1698134270432&to=1713945470432&panelId=1" width="450" height="200';
            $('#dashboard-iframe').attr('src', dashboardUrl);
            $('#dashboard-iframe').show();
            $('#bill-maintenance-content').hide(); // hide main content when iframe is shown
            $('#upload-content').hide(); // hide main content when iframe is shown
        } else if (this.children[0].id === 'bill-maintenance-link') {
            $('#upload-content').hide();
            $('#dashboard-iframe').hide();
            $('#bill-maintenance-content').show();
        } else {
            $('#dashboard-iframe').hide(); // hide iframe when main content is shown
            $('#bill-maintenance-content').hide();
            $('#upload-content').show();
        }

        // Simulate a delay for loading the new page
        setTimeout(function () {
            // Hide the loading animation
            $('#loading').hide();
        }, 500);
    });

    $('#upload-form').on('submit', function (e) {
        e.preventDefault();

        var formData = new FormData(this);
        var platform = $('#platform').val();
        var url = platform === 'weichat' ? '/load_weichat_bills' : '/load_alipay_bills';
        $.ajax({
            type: 'POST',
            url: url,
            data: formData,
            cache: false,
            contentType: false,
            processData: false,
            success: function (data) {
                alert('文件上传成功');
            },
            error: function (data) {
                alert('文件上传出错');
            }
        });
    });
});