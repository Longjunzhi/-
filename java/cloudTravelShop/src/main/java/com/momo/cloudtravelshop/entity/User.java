package com.momo.cloudtravelshop.entity;

import com.baomidou.mybatisplus.annotation.TableLogic;
import lombok.AllArgsConstructor;
import lombok.Data;
import lombok.NoArgsConstructor;

@Data
@AllArgsConstructor
@NoArgsConstructor
public class User {

    private Long id;

    private String accountName;

    private String mobile;

    private String password;

    @TableLogic
    private Integer isDeleted;

    private String createTime;

    private String updateTime;

}
